"use client";
import { useRouter } from "next/navigation.js";
import { useState } from "react";
import { BiWorld } from "react-icons/bi";
import { MdPrivacyTip } from "react-icons/md";
import { SiGnuprivacyguard } from "react-icons/si";
import { DisplayComments } from "../components/modals/displayComment";

const DisplayPost = ({ postData, idUser, onCommentClick }) => {
  const [postDataState, setPostDataState] = useState(postData),
    [isVisibleComment, setIsVisibleComment] = useState(false),
    [commentCounter, setCommentCounter] = useState(postDataState.Comments);
  const router = useRouter();
  const handleProfileClick = () => {
    router.push(`/home/profile?id=${postDataState.User_id}`);
  };
  return (
    <div className="w-full post_div mb-5 border border-gray-700 p-2 rounded-md">
      <div className="post_div_top flex flex-col gap-1 w-fit justify-center ">
        <div className="header_pos w-fit">
          <div className="info_post flex items-start gap-2">
            <img
              src={
                postDataState.Avatar != "NULL"
                  ? `${postDataState.Avatar}`
                  : "/assets/profilibg.jpg"
              }
              alt="Profile picture"
              onClick={handleProfileClick}
              className="profile_pic rounded-full cursor-pointer hover:opacity-60"
              width={50}
              height={50}
            />
            <div className="flex flex-col">
              <div className="flex gap-2 items-center w-[100%]">
                <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                  {postDataState.FullName}
                </h3>
                <span className="username_post italic text-primary">
                  @{postDataState.UserName}
                </span>
                <div className="flex">
                  {/* <i><small>{postDataState.Privacy}</small></i> */}
                  {(postDataState.Privacy == "private" && (
                    <SiGnuprivacyguard className="text-2xl" />
                  )) ||
                    (postDataState.Privacy == "public" && (
                      <BiWorld className="text-2xl" />
                    )) ||
                    (postDataState.Privacy == "almost" && (
                      <MdPrivacyTip className="text-2xl" />
                    ))}
                </div>
              </div>
              <div className="flex gap-4 text-sm">
                <span>{postDataState.Date}</span>
                <div className="flex gap-1 italic text-primary">
                  {postDataState.Categories.length != 0 &&
                    postDataState.Categories.split(",").map((tag) => (
                      <span key={postDataState.Id + tag} className="mr-1">
                        #{tag}
                      </span>
                    ))}
                </div>
              </div>
              <p className="title_post  break-words w-[100%] md:max-w-[300px] max-w-[300px] lg:max-w-[600px]">
                {postDataState.Content}
              </p>
            </div>
          </div>
        </div>
        {postDataState.Media != "NULL" && (
          <img
            src={postDataState.Media}
            alt="Post image"
            className="img_post max-w-full z-0 hover:shadow-xl overflow-hidden cursor-pointer transition duration-300 ease-linear scale-95 hover:scale-100"
            width={700}
            height={200}
          />
        )}
      </div>
      <div className="footer_post w-fit flex gap-4">
        <button
          className="comment_post flex gap-1 items-center"
          onClick={() => {
            setIsVisibleComment(true);
          }}
        >
          <img
            src="/assets/icons/comment.png"
            alt="Comment icon"
            width={30}
            height={30}
          />
          <span>{commentCounter}</span>
        </button>
      </div>
      <DisplayComments
        isVisible={isVisibleComment}
        key={postDataState.ID}
        postId={postDataState.ID}
        idUser={idUser}
        onClose={() => setIsVisibleComment(false)}
        increment={() => setCommentCounter((prevCount) => prevCount + 1)}
      />
    </div>
  );
};

export default DisplayPost;
