"use client";

import { WebSocketContext } from "@/app/_lib/websocket.js";
import {
  followUser,
  getProfileById,
  profileTypeUser,
  unfollowUser,
} from "@/app/api/api.js";
import { useSearchParams } from "next/navigation.js";
import { useContext, useState } from "react";
import { useQuery } from "react-query";
import DisplayPost from "../../home/displayPost.js";
import { ShowFollowers } from "../modals/showfollowers.js";
import { ShowFollowings } from "../modals/showfollowing.js";

export default function Profile({ sessionId, sessionToken }) {
  const [IsVisibleFollowers, setIsVisibleFollowers] = useState(false);
  const [IsVisibleFollowing, setIsVisibleFollowing] = useState(false);
  const [user, setUser] = useState({});
  const [followers, setFollowers] = useState([]);
  const [followings, setFollowings] = useState([]);
  const [posts, setPosts] = useState([]);
  const { sendMessage } = useContext(WebSocketContext);
  //recup id from url
  const searchParams = useSearchParams();
  let userid = searchParams.get("id");

  const id = userid ? parseInt(userid) : sessionId;
  useQuery("profile", () => getProfileById(id, sessionId, sessionToken), {
    enabled: true,
    refetchInterval: 5000,
    staleTime: 5000,
    onSuccess: (data) => {
      setUser(data?.user);
      setFollowers(data?.followers);
      setFollowings(data?.followings);
      setPosts(data?.posts);
    },
    onError: (error) => console.error("Query Profile error:", error),
  });
  const followersId = followers?.map((follower) => follower.id);
  const HandleFollow = (id, sessionId) => {
    if (sessionId && followersId?.includes(sessionId)) {
      try {
        const response = unfollowUser(id, sessionId, sessionToken);
        if (response.error) {
          console.error("error unfollowing");
        } else {
          console.error("unfollowed");
        }
      } catch (error) {
        console.error("error unfollowing", error);
      }
    } else {
      if (user?.user_type == "public") {
        try {
          const response = followUser(id, sessionId, sessionToken);
          if (response.error) {
            console.error("error following");
          } else {
            console.error("followed");
          }
        } catch (error) {
          console.error("error following", error);
        }
      } else {
        sendMessage({
          type: "follow",
          followed_id: id,
          follower_id: sessionId,
        });
      }
    }
  };

  const SwitchTypeProfile = (sessionId, user_type) => {
    try {
      const response = profileTypeUser(sessionId, user_type, sessionToken);
      if (response.error) {
        console.error("error switching", response.error);
      }
    } catch (error) {
      console.error("error switching", error);
    }
  };

  // map followers to recup id in a array

  const isPublic = user?.user_type === "public";
  const isOwner = sessionId == id;
  const isFollower = sessionId && followersId?.includes(sessionId);

  if (!user) {
    return <div>Loading...</div>;
  }
  return (
    <div className="md:w-[300px] lg:w-[500px] xl:w-[700px] 2xl:w-[1000px] w-screen h-full flex flex-col gap-2">
      <div>
        <div className="w-full h-24 bg-black bg-cover bg-center bg-no-repeat"></div>
        <div className="p-4">
          <div className="relative flex w-full">
            <div className="flex flex-1">
              <div className="h-36 w-36 -mt-24 relative rounded-full">
                <img
                  className="h-36 w-36 rounded-full border-4 border-gray-900"
                  src={
                    user.avatar !== "NULL"
                      ? user.avatar
                      : "/assets/profilibg.jpg"
                  }
                  alt={user?.first_name + " " + user?.last_name}
                />
              </div>
            </div>
            <div className="flex flex-col text-right">
              {isOwner ? (
                <button
                  className="ml-auto mr-0 flex max-h-max max-w-max items-center justify-center whitespace-nowrap rounded-full border border-blue-500 bg-transparent px-4 py-2 font-bold text-blue-500 hover:border-blue-800 hover:shadow-lg focus:outline-none focus:ring"
                  onClick={() => SwitchTypeProfile(sessionId, user?.user_type)}
                >
                  {!isPublic ? "Switch to Public" : "Switch to Private"}
                </button>
              ) : (
                <button
                  className="ml-auto mr-0 flex max-h-max max-w-max items-center justify-center whitespace-nowrap rounded-full border border-blue-500 bg-transparent px-4 py-2 font-bold text-blue-500 hover:border-blue-800 hover:shadow-lg focus:outline-none focus:ring"
                  onClick={() => HandleFollow(id, sessionId)}
                >
                  {isFollower ? "Unfollow" : "Follow"}
                </button>
              )}
            </div>
          </div>
          <div className="ml-3 mt-3 w-full justify-center space-y-1">
            {/* User basic*/}
            <div>
              <h2 className="text-xl font-bold leading-6 text-white">
                {user?.first_name} {user?.last_name}
              </h2>
              <p className="text-sm font-medium leading-5 text-gray-600">
                {user?.user_name ? "@" + user?.user_name : ""}
              </p>
              <p className="text-sm pt-1 leading-tight text-white">
                {user?.gender}
              </p>
            </div>
            {/* Description and others */}
            {isPublic || isOwner || isFollower ? (
              <>
                <div className="pt-1">
                  <p className="mb-2 leading-tight text-white">
                    {user?.about_me}
                  </p>
                  <div className="flex text-gray-600">
                    <span className="mr-2 flex">
                      <svg viewBox="0 0 24 24" className="paint-icon h-5 w-5">
                        <g>
                          <path d="M11.96 14.945c-.067 0-.136-.01-.203-.027-1.13-.318-2.097-.986-2.795-1.932-.832-1.125-1.176-2.508-.968-3.893s.942-2.605 2.068-3.438l3.53-2.608c2.322-1.716 5.61-1.224 7.33 1.1.83 1.127 1.175 2.51.967 3.895s-.943 2.605-2.07 3.438l-1.48 1.094c-.333.246-.804.175-1.05-.158-.246-.334-.176-.804.158-1.05l1.48-1.095c.803-.592 1.327-1.463 1.476-2.45.148-.988-.098-1.975-.69-2.778-1.225-1.656-3.572-2.01-5.23-.784l-3.53 2.608c-.802.593-1.326 1.464-1.475 2.45-.15.99.097 1.975.69 2.778.498.675 1.187 1.15 1.992 1.377.4.114.633.528.52.928-.092.33-.394.547-.722.547z" />
                          <path d="M7.27 22.054c-1.61 0-3.197-.735-4.225-2.125-.832-1.127-1.176-2.51-.968-3.894s.943-2.605 2.07-3.438l1.478-1.094c.334-.245.805-.175 1.05.158s.177.804-.157 1.05l-1.48 1.095c-.803.593-1.326 1.464-1.475 2.45-.148.99.097 1.975.69 2.778 1.225 1.657 3.57 2.01 5.23.785l3.528-2.608c1.658-1.225 2.01-3.57.785-5.23-.498-.674-1.187-1.15-1.992-1.376-.4-.113-.633-.527-.52-.927.112-.4.528-.63.926-.522 1.13.318 2.096.986 2.794 1.932 1.717 2.324 1.224 5.612-1.1 7.33l-3.53 2.608c-.933.693-2.023 1.026-3.105 1.026z" />
                        </g>
                      </svg>
                      <span className="ml-1 leading-5 text-blue-400">
                        {user?.email}
                      </span>
                    </span>
                    <span className="mr-2 flex">
                      <svg viewBox="0 0 24 24" className="paint-icon h-5 w-5">
                        <g>
                          <path d="M19.708 2H4.292C3.028 2 2 3.028 2 4.292v15.416C2 20.972 3.028 22 4.292 22h15.416C20.972 22 22 20.972 22 19.708V4.292C22 3.028 20.972 2 19.708 2zm.792 17.708c0 .437-.355.792-.792.792H4.292c-.437 0-.792-.355-.792-.792V6.418c0-.437.354-.79.79-.792h15.42c.436 0 .79.355.79.79V19.71z" />
                          <circle cx="7.032" cy="8.75" r="1.285" />
                          <circle cx="7.032" cy="13.156" r="1.285" />
                          <circle cx="16.968" cy="8.75" r="1.285" />
                          <circle cx="16.968" cy="13.156" r="1.285" />
                          <circle cx={12} cy="8.75" r="1.285" />
                          <circle cx={12} cy="13.156" r="1.285" />
                          <circle cx="7.032" cy="17.486" r="1.285" />
                          <circle cx={12} cy="17.486" r="1.285" />
                        </g>
                      </svg>
                      <span className="ml-1 leading-5">
                        Born {formatDate(user?.birth_date)}
                      </span>
                    </span>
                  </div>
                </div>
                {/*Each section shows the number of followers or followings along with their avatars. */}
                <div className="flex w-full items-start justify-start divide-x divide-solid divide-white py-4">
                  <div
                    className="flex flex-rows gap-x-1 pr-3 text-center cursor-pointer hover:text-primary "
                    onClick={() => {
                      setIsVisibleFollowing(true);
                    }}
                  >
                    {followings ? (
                      <span className="flex -space-x-1 mr-1">
                        {followings.map((following) => (
                          <img
                            key={following.first_name + following.last_name}
                            className="inline-block h-6 w-6 rounded-full ring-2 ring-white"
                            src={
                              following.avatar !== "NULL"
                                ? following.avatar
                                : "/assets/profilibg.jpg"
                            }
                            alt={`${following.first_name} ${following.last_name}`}
                          />
                        ))}
                      </span>
                    ) : (
                      ""
                    )}
                    <span className="font-bold text-white">
                      {followings ? followings.length : 0}
                    </span>
                    <span className="text-gray-600">
                      Following{followings?.length ? "s" : ""}
                    </span>
                  </div>
                  <div
                    className="flex flex-rows gap-x-1 px-3 text-center cursor-pointer"
                    onClick={() => {
                      setIsVisibleFollowers(true);
                    }}
                  >
                    {followers ? (
                      <span className="flex -space-x-1 mr-1">
                        {followers.map((follower) => (
                          <img
                            key={follower.id}
                            className="inline-block h-6 w-6 rounded-full ring-2 ring-white"
                            src={
                              follower.avatar !== "NULL"
                                ? follower.avatar
                                : "/assets/profilibg.jpg"
                            }
                            alt={`${follower.first_name} ${follower.last_name}`}
                          />
                        ))}
                      </span>
                    ) : (
                      ""
                    )}
                    <span className="font-bold text-white">
                      {followers ? followers.length : 0}
                    </span>
                    <span className="text-gray-600">
                      Follower{followers?.length ? "s" : ""}
                    </span>
                  </div>
                </div>
              </>
            ) : (
              ""
            )}
          </div>
        </div>
        <hr />
      </div>
      {isPublic || isOwner || isFollower ? (
        <div className="post_div_top flex flex-col items-center">
          {posts?.map((post) => (
            <DisplayPost key={post.ID} postData={post} />
          ))}
        </div>
      ) : (
        ""
      )}
      <ShowFollowers
        followers={followers}
        isVisible={IsVisibleFollowers}
        onClose={() => {
          setIsVisibleFollowers(false);
        }}
      />
      <ShowFollowings
        followings={followings}
        isVisible={IsVisibleFollowing}
        onClose={() => {
          setIsVisibleFollowing(false);
        }}
      />
    </div>
  );
}

function formatDate(dateString) {
  const options = { year: "numeric", month: "long", day: "numeric" };
  const date = new Date(dateString);
  return date.toLocaleDateString("fr-FR", options);
}
