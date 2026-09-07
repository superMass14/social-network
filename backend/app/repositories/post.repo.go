package repositories

import (
	db "backend/database"
	"backend/models"
	"backend/utils"
	"errors"
	"log"
	"strings"

	"github.com/gofrs/uuid"
)

type PostRepository struct {
	BaseRepo
}

const (
	GetPostQuery = `
SELECT
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
    p.user_id AS post_user_id,
	p.privacy,
	u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(DISTINCT c.id) AS comment_count,
	GROUP_CONCAT(DISTINCT cat.category) AS categories
FROM
    posts AS p
LEFT JOIN
    comment AS c ON p.id = c.post_id,
	categories AS cat ON p.id = cat.post_id,
	users AS u ON p.user_id = u.id
WHERE
    (
        p.privacy = 'public'
        OR (
            p.privacy = 'private' AND (
                p.user_id = ? -- Post creator
                OR EXISTS (
                    SELECT 1
                    FROM users_followers
                    WHERE user_id_followed = p.user_id
                    AND user_id_follower = ?
                )
                OR p.user_id = ? --? is the user_id of the logged in user

            )
        )
        OR (
            p.privacy = 'almost' AND EXISTS (
                SELECT 1
                FROM viewers
                WHERE user_id = ?
                AND post_id = p.id
            )
        )
    )
	AND p.group_id = 0
	GROUP BY p.id, p.content, p.media, p.date, p.user_id
	ORDER BY p.timestamp DESC;
`
	GetPostGroupQuery = `
SELECT
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
	u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(DISTINCT c.id) AS comment_count
FROM
    posts AS p
LEFT JOIN
    comment AS c ON p.id = c.post_id,
	users AS u ON p.user_id = u.id
WHERE group_id = ?
	GROUP BY p.id, p.content, p.media, p.date, p.user_id
	ORDER BY p.timestamp DESC;
`
	GetProfilePost = `
SELECT
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
    p.user_id AS post_user_id,
    p.privacy,
    u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(DISTINCT c.id) AS comment_count,
    GROUP_CONCAT(DISTINCT cat.category) AS categories
FROM
    posts AS p
    LEFT JOIN comment AS c ON p.id = c.post_id,
    categories AS cat ON p.id = cat.post_id,
    users AS u ON  p.user_id=u.id
WHERE
p.user_id = ?
AND
    (
        p.privacy = 'public'
        OR (
            p.privacy = 'private'
            AND (
                p.user_id = ? --? Post creator
                OR EXISTS (
                    SELECT
                        1
                    FROM
                        users_followers
                    WHERE
                        user_id_followed = p.user_id
                        AND user_id_follower = ? --?
                )
				OR p.user_id = ? --? id of logged user
            )
        )
        OR (
            p.privacy = 'almost'
            AND EXISTS (
                SELECT
                    1
                FROM
                    viewers
                WHERE
                    user_id = ? --?
                    AND post_id = p.id
            )
        )

    )
    AND p.group_id =0
GROUP BY
    p.id,
    p.content,
    p.media,
    p.date,
    p.user_id
	ORDER BY p.timestamp DESC;`
)

func (p *PostRepository) init() {
	p.DB = db.DB
	p.TableName = "posts"
}

func (p *PostRepository) CreatePost(post models.Post) (bool, models.Errormessage) {

	//todo: checking Id_user validity

	//checking Title's validity
	log.Println("ins content => ", post.ToIns.Content, " length => ", len([]rune(post.ToIns.Content)))
	log.Println("ins media => ", post.ToIns.Media, " length => ", len([]rune(post.ToIns.Media)))

	if strings.TrimSpace(post.ToIns.Content) == "" && strings.TrimSpace(post.ToIns.Media) == "NULL" {
		log.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to empty content and media ❌\n", "1")
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to empty input",
				StatusCode: models.BRstatus,
				Display:    false,
			}
	}

	//checking category's validity
	if len(post.Categories) < 1 { //user did not select a categorie
		log.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to missing category❌\n", "1")
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to missing category",
				StatusCode: models.BRstatus,
				Display:    false,
			}
	}

	if len(post.ToIns.Content) > 1500 { //found only spaces,newlines in the input or chars number limit exceeded
		log.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to invalid input ❌\n", "1")
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to invalid input",
				StatusCode: models.BRstatus,
				Display:    false,
			}
	}

	errIns := p.InsertPost(post)
	if errIns != nil {
		return true,
			models.Errormessage{
				Type:       models.ISEtype,
				Msg:        models.ISEmsg,
				StatusCode: models.ISEstatus,
				Display:    false,
			}
	}
	return false, models.Errormessage{}
}

func (P *PostRepository) InsertPost(post models.Post) error {
	id_post, errp := uuid.NewV4()
	if errp != nil {
		log.Println("❌ Create_post ⚠ ERROR ⚠ : couldn't generate a unique post id")
		return errp
	}

	date, time := utils.Time() //date and time
	post.ToIns.Date = date + " " + time
	// inserting value in database
	//-- formatting value's special chars
	if post.ToIns.Content != "" {
		post.ToIns.Content = utils.EncodeValue(post.ToIns.Content)
	}
	post.ToIns.ID = id_post.String()
	err := P.DB.Insert(P.TableName, post.ToIns)
	if err != nil {
		log.Println("❌ error while inserting post", err)
		return err
	}
	//inserting categories
	formatedCatego := utils.FormatCategory(post.Categories, id_post.String())
	for i := range formatedCatego {
		errCat := P.DB.Insert("categories", formatedCatego[i])
		if errCat != nil {
			log.Println("❌ error while inserting categories")
			return err
		}
	}
	if post.ToIns.Privacy == "almost" {
		formatedViewers := utils.FormatViewers(post.Viewers, id_post.String())
		for i := range formatedViewers {
			errView := P.DB.Insert("viewers", formatedViewers[i])
			if errView != nil {
				log.Println("❌ error while inserting viewers")
				return err
			}
		}
	}
	log.Println("✅ post has been created successfully")
	return nil
}

func (P *PostRepository) LoadPost(IdUser int) ([]models.DataPost, error) {
	var postTab []models.DataPost

	rows, err := P.DB.Query(GetPostQuery, IdUser, IdUser, IdUser, IdUser)
	if err != nil {
		log.Println("❌ Error while retrieving posts => ", err)
		return nil, errors.New("error while retrieving posts from the database")
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.DataPost
		errScan := rows.Scan(&temp.ID, &temp.Content, &temp.Media, &temp.Date, &temp.User_id, &temp.Privacy, &temp.Avatar, &temp.UserName, &temp.FullName, &temp.Comments, &temp.Categories)
		if errScan != nil {
			log.Println("⚠ GetPost scan err ⚠ :", errScan)
			return nil, errors.New("error while scanning")
		}

		temp.Content = utils.DecodeValue(temp.Content)
		postTab = append(postTab, temp)
	}
	return postTab, nil
}

func (p *PostRepository) LoadPostGroup(GroupID int) ([]models.DataPost, error) {
	var postTab []models.DataPost
	rows, err := p.DB.Query(GetPostGroupQuery, GroupID)
	if err != nil {
		log.Println("❌ Error while retrieving in GroupPost => ", err)
		return nil, errors.New("error while retrieving onepost from the database")
	}
	defer rows.Close()

	var data models.DataPost
	//! modify retrieval
	for rows.Next() {
		errScan := rows.Scan(&data.ID, &data.Content, &data.Media, &data.Date, &data.Avatar, &data.UserName, &data.FullName, &data.Comments)
		if errScan != nil {
			log.Println("⚠ loadPostGroup scan err ⚠ :", errScan)
			return nil, errors.New("error while scanning")
		}
		data.Content = utils.DecodeValue(data.Content)
		postTab = append(postTab, data)
	}
	return postTab, nil
}

func (p *PostRepository) LoadPostGroupByUserID(IdUser int, IdProfil int) ([]models.DataPost, error) {
	var postTab []models.DataPost
	rows, err := p.DB.Query(GetProfilePost, IdProfil, IdProfil, IdUser, IdUser, IdUser)
	if err != nil {
		log.Println("❌ Error while retrieving posts => ", err)
		return nil, errors.New("error while retrieving posts from the database")
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.DataPost
		errScan := rows.Scan(&temp.ID, &temp.Content, &temp.Media, &temp.Date, &temp.User_id, &temp.Privacy, &temp.Avatar, &temp.UserName, &temp.FullName, &temp.Comments, &temp.Categories)
		if errScan != nil {
			log.Println("⚠ GetPost scan err profile ⚠ :", errScan)
			return nil, errors.New("error while scanning")
		}

		temp.Content = utils.DecodeValue(temp.Content)
		postTab = append(postTab, temp)
	}
	return postTab, nil
}
