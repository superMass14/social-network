"use client";
import { useState } from "react";
import { useQuery } from "react-query";
import { PostForm } from "./postForm";
import DisplayPost from "./displayPost";

const HomePage = ({ id }) => {
  const [formCreatePost, setFormCreatePost] = useState(false);
  const [followers, setFollowers] = useState([]);

  const [posts, setPosts] = useState([]),
    [fetchState, setFetchState] = useState(true),
    userID = new FormData();
  userID.append("userID", id);
  userID.append("type", "loadPosts");
  const options = {
    method: "POST",
    body: userID,
  };
  const fetchPosts = async () => {
    const datas = new FormData();
    // const user = await getSessionUser();
    // datas["userID"] = user.id;
    try {
      const response = await fetch("http://localhost:8080/post", options);
      const data = await response.json();
      return {
        posts: data.data,
        followers: data.followers,
        errType: data.status,
        errMess: data.msg,
      };
    } catch (error) {
      console.error("Error while querying posts ", error);
      setFetchState(false);
      return Promise.reject(error);
    }
  };

  useQuery("posts", fetchPosts, {
    enabled: fetchState,
    refetchInterval: 1000,
    staleTime: 500,
    onSuccess: (newData) => {
      //  if (newData.errType == 400) setFetchState(false);
      setPosts(newData.posts);
      setFollowers(newData.followers);
    },
    onError: (error) => {
      setFetchState(false);
      console.error("Query error in posts:", error);
    },
  });
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen ">
      <div className="flex justify-center mb-4">
        {PostForm(
          id,
          formCreatePost,
          () => setFormCreatePost(false),
          followers
        )}
      </div>

      <button
        className="inline-flex items-center px-4 py-4 z-10 text-m font-semibold absolute lg:end-72 md:end-52 end-4 md:bottom-2 bottom-20 text-center text-white bg-primary rounded-lg hover:bg-second"
        onClick={() => {
          setFormCreatePost(!formCreatePost);
        }}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={2}
          stroke="currentColor"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M12 4.5v15m7.5-7.5h-15"
          />
        </svg>
        <p className="lg:block hidden">NEW POST</p>
      </button>

      <div className="post_div_top flex flex-col items-center">
        {posts?.map((e) => (
          <DisplayPost
            key={e.ID}
            postData={e}
            idUser={id}
            onLikeClick={onLikeClick}
            onDislikeClick={onDislikeClick}
            onCommentClick={onCommentClick}
          />
        ))}
      </div>
    </div>
  );
};

export default HomePage;

const onLikeClick = () => {
  alert("like");
};

const onDislikeClick = () => {
  alert("dislike");
};

const onCommentClick = () => {
  alert("Comment disp");
};

const onProfileClick = () => {
  alert("profile disp");
};

// let postData = {
//   id: 3,
//   content: "This is the content of the third post.",
//   media: "imagepost.jpg",
//   date: "2024-03-05",
//   userId: 1,
//   fullname: "Madike Yade",
//   username: "dickss",
//   user: {
//     id: 1,
//     first_name: "Madike",
//     last_name: "Yade",
//     user_name: "dickss",
//     gender: "Male",
//     user_type: "private",
//     birth_date: "2000-01-01",
//     avatar: "profilibg.jpg",
//     about_me: "about me...",
//     password: "1234",
//     email: "dickss@gmail.com",
//   },
//   groupId: 0,
//   privacy: "almost",
// };
