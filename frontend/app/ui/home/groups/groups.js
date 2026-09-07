"use client";
import { WebSocketContext } from "@/app/_lib/websocket";
import { useSession } from "@/app/api/api";
import Link from "next/link";
import { useContext, useState } from "react";
import { useQuery } from "react-query";
import { CreateGroup } from "../../components/modals/createGroup";
import { JoinGroup } from "./group.utils/joinGroup";

const Groups = () => {
  const { session, errSess } = useSession();
  const [formCreateGr, setFormCreateGr] = useState(false);
  const [groupData, setGroupData] = useState([]);
  const [fetcherState, setFetcherState] = useState(true);
  const [groupJoined, setGroupJoined] = useState([]);

  const fetchGroups = async () => {
    try {
      let token = localStorage.getItem("token");
      // console.log(token);
      const response = await fetch(
        `http://localhost:8080/groups?token=${token}`
      );
      const data = await response.json();
      return { groupJoined: data[0], groupData: data[1] };
    } catch (error) {
      console.error("Erreur ", error);
      return Promise.reject(error);
    }
  };

  useQuery("groups", fetchGroups, {
    enabled: fetcherState,
    refetchInterval: 5000,
    staleTime: 1000,
    onSuccess: (newData) => {
      setGroupJoined(newData?.groupJoined);
      setGroupData(newData?.groupData);
    },
    onError: (error) => {
      console.error("Query error:", error);
    },
  });

  return (
    <div className="md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1200px] w-screen h-full flex flex-col">
      {/* {isLoading ? (
                <div>
                    chargement ..
                </div>
            ) : ( */}
      <>
        <div className="w-[90%] justify-between flex items-center mb-5 gap-5">
          <h1 className="text-xl font-bold">Communities</h1>
          <button
            className="inline-flex items-center px-4 py-2 text-m font-semibold text-center text-white bg-primary rounded-lg hover:bg-second"
            onClick={() => {
              setFormCreateGr(true);
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
            New Group
          </button>
        </div>
        <div className="w-full flex gap-3 overflow-x-scroll">
          {groupJoined?.map((group) => (
            <GroupCard key={group.id} isMember={true} {...group} />
          ))}
        </div>
        <h1 className="text-xl font-bold my-5">Discover new communities</h1>
        <div className="w-full flex gap-3 overflow-x-scroll pb-10">
          {groupData
            ? groupData.map((group) => (
                <GroupCard key={group.id} isMember={false} {...group} />
              ))
            : ""}
        </div>
        <CreateGroup
          isVisible={formCreateGr}
          onClose={() => {
            setFormCreateGr(false);
          }}
        />
      </>
      {/* )} */}
    </div>
  );
};

export default Groups;

const GroupCard = ({ isMember, id, image, name, description, href, state }) => {
  if (description.length > 50) {
    description = description.slice(0, 50) + " ...";
  }

  // const { sendMessage, readMessages, messages } = useApi();

  const { sendMessage, readMessages, messages } = useContext(WebSocketContext);

  const handleLoginJoinMessage = () => {
    const joinMessage = messages.find(
      (message) => message.Type === "JoinGroup"
    );
    if (joinMessage) {
      messages.map((message) => {
        if (message.Data.id_group === id) {
          return;
        }
      });
    }
  };
  // useEffect(() => {
  handleLoginJoinMessage();
  // }, [sendMessage]);

  return (
    <>
      {isMember ? (
        <Link
          href={href}
          className="inline-flex items-center text-m font-semibold text-center text-white rounded-l"
        >
          <div className="w-[200px] border rounded-lg shadow-xl min-h-[206px] bg-bg bg-clip-padding backdrop-filter backdrop-blur-sm bg-opacity-5 border-gray-800 hover:bg-opacity-5 hover:bg-primary cursor-pointer">
            <div className="flex flex-col items-center py-3">
              <img
                src={image}
                alt={name}
                width={500}
                height={500}
                className="w-24 h-24 mb-3 rounded-full hover:scale-110 transition duration-200 ease-in shadow-lg"
              />
              <h5 className="mb-1 text-xl font-medium text-white text-center">
                {name}
              </h5>
              <span className="max-h-14 overflow-hidden text-sm text-gray-300 text-center">
                {description}
              </span>
            </div>
          </div>
        </Link>
      ) : (
        <div className="w-[200px] border rounded-lg shadow-xl bg-primary bg-opacity-0 bg-clip-padding backdrop-filter backdrop-blur-md hover:bg-opacity-5 hover:bg-primary border-gray-800 cursor-pointer">
          <div className="flex flex-col items-center h-[100%]  justify-between py-3">
            {image ? (
              <img
                src={`${image}`}
                alt={name}
                width={200}
                height={200}
                className="w-24 h-24 mb-3 hover:scale-110 duration-200 transition ease-in-out rounded-full shadow-lg"
              />
            ) : (
              ""
            )}
            <h5 className="mb-1 text-xl font-medium text-white text-center">
              {name}
            </h5>
            <span className="max-h-14 overflow-hidden text-sm text-gray-300 text-center">
              {description ? description : ""}
            </span>

            <div className="flex mt-4 md:mt-6">
              {state !== "disable" ? (
                <button
                  onClick={() => {
                    JoinGroup(id, sendMessage, readMessages);
                  }}
                  className="inline-flex items-center px-4 py-2 text-m font-semibold text-center text-white bg-primary rounded-lg hover:bg-second"
                >
                  Join
                </button>
              ) : (
                ""
              )}
            </div>
          </div>
        </div>
      )}
    </>
  );
};
