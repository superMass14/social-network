"use client";
import { CreateEvent } from "@/app/ui/components/modals/createEvent";
import { CreatePost } from "@/app/ui/components/modals/createPost";
import { DisplayMembers } from "@/app/ui/components/modals/displayMembers";
import { Suggest } from "@/app/ui/components/modals/suggest";
import { usePathname } from "next/navigation";
import { useContext, useState } from "react";
import { useQuery } from "react-query";
import DisplayPost from "../../displayPost";

// import { useApi } from "@/app/_lib/utils";
import { WebSocketContext } from "@/app/_lib/websocket";
import { displayEventToJoin } from "./displayEventToJoin";
import { displayEvents } from "./displayEvents";

const CardEvent = ({ description, date }) => {
  return (
    <Card sx={{ maxWidth: 400, mb: 2 }}>
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: "teal" }}>{creatorName.charAt(0)}</Avatar>
        }
        title={creatorName}
        subheader={`${date} à ${time}`}
      />
      <CardContent>
        <Typography variant="body2" color="textSecondary">
          {description}
        </Typography>
      </CardContent>
    </Card>
  );
};

const Group = ({ sessionID }) => {
  const { sendMessage } = useContext(WebSocketContext);

  const [formCreateEv, setFormCreateEv] = useState(false);
  const [formCreateP, setFormCreateP] = useState(false);
  const [groupPosts, setGroupPosts] = useState();
  const [groupInfo, setGroupInfo] = useState({});
  const [suggestFriend, setSuggestFriend] = useState(false);
  const [isVisibleMembers, setIsVisibleMembers] = useState(false);
  const [members, setMembers] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [events, setEvents] = useState([]);
  const [eventsJoined, setEventsJoined] = useState([]);

  const pathname = usePathname();
  const id = pathname.split("id=")[1];

  const fetchGroups = async () => {
    const token = localStorage.getItem("token") || null;

    try {
      const url = `http://localhost:8080/groups/group?id=${id}&token=${token}`;

      const response = await fetch(url, {
        method: "GET",
      });
      const data = await response.json();
      return data;
    } catch (error) {
      console.error("Erreur ", error);
      return Promise.reject(error);
    }
  };

  useQuery("groups", fetchGroups, {
    enabled: true,
    refetchInterval: 2000,
    staleTime: 1000,
    onSuccess: (newData) => {
      setGroupPosts(newData.posts);
      setGroupInfo(newData.group_data);
      setMembers(newData.members);
      setFollowers(newData.followers);
      console.log(followers);
      setEvents(newData.events);
      setEventsJoined(newData.events_joined);
    },
    onError: (error) => {
      console.error("Query error:", error);
    },
  });

  return (
    <div
      className="md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1100px] w-screen h-full
                        flex flex-col gap-2"
    >
      <div className="w-full h-60 mb-3">
        {groupInfo?.image ? (
          <img
            src={`${groupInfo.image}`}
            alt="cover"
            width={1000}
            height={1000}
            className="w-full max-h-[250px] object-cover scale-100 hover:scale-105  rounded-sm  transition duration-300 ease-in shadow-lg"
          />
        ) : (
          ""
        )}
      </div>
      <div className="flex items-center justify-between 2xl:w-[90%] w-[95%]">
        <div className="">
          <h1 className="text-2xl font-bold underline underline-offset-4 mb-2 flex items-center gap-2 ">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M4.5 3.75a3 3 0 0 0-3 3v10.5a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3V6.75a3 3 0 0 0-3-3h-15Zm4.125 3a2.25 2.25 0 1 0 0 4.5 2.25 2.25 0 0 0 0-4.5Zm-3.873 8.703a4.126 4.126 0 0 1 7.746 0 .75.75 0 0 1-.351.92 7.47 7.47 0 0 1-3.522.877 7.47 7.47 0 0 1-3.522-.877.75.75 0 0 1-.351-.92ZM15 8.25a.75.75 0 0 0 0 1.5h3.75a.75.75 0 0 0 0-1.5H15ZM14.25 12a.75.75 0 0 1 .75-.75h3.75a.75.75 0 0 1 0 1.5H15a.75.75 0 0 1-.75-.75Zm.75 2.25a.75.75 0 0 0 0 1.5h3.75a.75.75 0 0 0 0-1.5H15Z"
                clipRule="evenodd"
              />
            </svg>
            {groupInfo?.name}
          </h1>
          <p className="text-lg flex items-center gap-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12Zm8.706-1.442c1.146-.573 2.437.463 2.126 1.706l-.709 2.836.042-.02a.75.75 0 0 1 .67 1.34l-.04.022c-1.147.573-2.438-.463-2.127-1.706l.71-2.836-.042.02a.75.75 0 1 1-.671-1.34l.041-.022ZM12 9a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z"
                clipRule="evenodd"
              />
            </svg>
            {groupInfo?.description}
          </p>
        </div>
        <div className="flex gap-2 lg:flex-row flex-col">
          <button
            onClick={() => {
              setFormCreateP(true);
            }}
            className="inline-flex items-center px-1 py-1 text-sm font-bold text-center max-h-[50px]  bg-gray-700 border border-gray-500 rounded-lg hover:bg-opacity-70 hover:bg-primary"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25ZM12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25V9Z"
                clipRule="evenodd"
              />
            </svg>
            Post
          </button>
          <button
            onClick={() => {
              setFormCreateEv(true);
            }}
            className="inline-flex items-center px-1 py-1 text-sm font-bold border border-gray-500 bg-gray-800 text-center max-h-[50px] rounded-lg hover:bg-opacity-70 hover:bg-primary"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path d="M12.75 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM7.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM8.25 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM9.75 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM10.5 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM12.75 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM14.25 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 13.5a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" />
              <path
                fillRule="evenodd"
                d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75Zm13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5Z"
                clipRule="evenodd"
              />
            </svg>
            Event
          </button>
          <button
            onClick={() => {
              setSuggestFriend(true);
            }}
            className="inline-flex items-center px-1 py-1 text-sm font-bold text-center max-h-[50px]  bg-gray-900 border border-gray-500 rounded-lg hover:bg-opacity-70 hover:bg-primary"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path d="M5.25 6.375a4.125 4.125 0 1 1 8.25 0 4.125 4.125 0 0 1-8.25 0ZM2.25 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63 13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63l-.001-.122ZM18.75 7.5a.75.75 0 0 0-1.5 0v2.25H15a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H21a.75.75 0 0 0 0-1.5h-2.25V7.5Z" />
            </svg>
            Suggest
          </button>
        </div>
      </div>

      <div className="flex items-center justify-between 2xl:w-[90%] w-[95%]">
        <p className="font-semibold flex items-center gap-1">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-6 h-6"
          >
            <path
              fillRule="evenodd"
              d="M18.685 19.097A9.723 9.723 0 0 0 21.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 0 0 3.065 7.097A9.716 9.716 0 0 0 12 21.75a9.716 9.716 0 0 0 6.685-2.653Zm-12.54-1.285A7.486 7.486 0 0 1 12 15a7.486 7.486 0 0 1 5.855 2.812A8.224 8.224 0 0 1 12 20.25a8.224 8.224 0 0 1-5.855-2.438ZM15.75 9a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z"
              clipRule="evenodd"
            />
          </svg>

          <span className="italic">Created by: </span>
          {groupInfo?.creator
            ? groupInfo?.creator.first_name + " " + groupInfo?.creator.last_name
            : ""}
        </p>
        <p
          className=" flex items-center gap-2 bg-primary bg-opacity-50 w-fit hover:bg-opacity-70 hover:text-white font-bold rounded-md cursor-pointer py-1 px-3 border-gray-700"
          onClick={() => {
            setIsVisibleMembers(true);
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-6 h-6"
          >
            <path
              fillRule="evenodd"
              d="M8.25 6.75a3.75 3.75 0 1 1 7.5 0 3.75 3.75 0 0 1-7.5 0ZM15.75 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM2.25 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM6.31 15.117A6.745 6.745 0 0 1 12 12a6.745 6.745 0 0 1 6.709 7.498.75.75 0 0 1-.372.568A12.696 12.696 0 0 1 12 21.75c-2.305 0-4.47-.612-6.337-1.684a.75.75 0 0 1-.372-.568 6.787 6.787 0 0 1 1.019-4.38Z"
              clipRule="evenodd"
            />
            <path d="M5.082 14.254a8.287 8.287 0 0 0-1.308 5.135 9.687 9.687 0 0 1-1.764-.44l-.115-.04a.563.563 0 0 1-.373-.487l-.01-.121a3.75 3.75 0 0 1 3.57-4.047ZM20.226 19.389a8.287 8.287 0 0 0-1.308-5.135 3.75 3.75 0 0 1 3.57 4.047l-.01.121a.563.563 0 0 1-.373.486l-.115.04c-.567.2-1.156.349-1.764.441Z" />
          </svg>
          Members:{" "}
          <span className="italic">{members ? members.length : "0"}</span>
        </p>
      </div>
      <div>
        <hr className="h-px m-2 bg-gray-700 border-0"></hr>
      </div>
      <div className="w-full flex justify-between">
        <div className="w-[20%]">
          <h1 className=" text-xl flex items-center gap-2 font-extrabold w-fit underline shadow-lg px-2 rounded-md ">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5"
              />
            </svg>
            Joined Events
          </h1>

          {eventsJoined ? displayEvents(eventsJoined) : ""}
          <h1 className=" text-xl flex items-center gap-2 font-extrabold w-fit underline shadow-lg px-2 rounded-md mt-4 ">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
            New Events
          </h1>
          {events ? displayEventToJoin(events, sendMessage) : ""}
        </div>
        <div className="w-[75%] ">
          {groupPosts
            ? groupPosts?.map((post) => {
                return (
                  <DisplayPost
                    key={post.ID}
                    postData={post}
                    idUser={sessionID}
                    onCommentClick={onCommentClick}
                  />
                );
              })
            : ""}
        </div>
      </div>
      <CreateEvent
        isVisible={formCreateEv}
        onClose={() => setFormCreateEv(false)}
        id={id}
        user={sessionID}
      />
      <CreatePost
        isVisible={formCreateP}
        onClose={() => setFormCreateP(false)}
        id={id}
        user={sessionID}
      />
      <Suggest
        followers={followers}
        isVisible={suggestFriend}
        onClose={() => setSuggestFriend(false)}
        id_group={id}
        sendMessage={sendMessage}
      />
      <DisplayMembers
        members={members}
        isVisible={isVisibleMembers}
        onClose={() => setIsVisibleMembers(false)}
      />
    </div>
  );
};

export default Group;

const onCommentClick = () => {
  alert("Comment disp");
};
