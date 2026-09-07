"use client";
// import { useApi } from "@/app/_lib/utils";
import { WebSocketContext } from "@/app/_lib/websocket";
import { useContext, useState } from "react";
// import { FaP } from "react-icons/fa6";
import { useQuery } from "react-query";

const Notification = () => {
  const { sendMessage, readMessages, messages } = useContext(WebSocketContext);
  const [notificationData, setNotificationData] = useState([]);

  const fetchNotifications = async () => {
    try {
      let token = localStorage.getItem("token");
      const response = await fetch(
        `http://localhost:8080/notifications?token=${token}`
      );
      const data = await response.json();
      return data;
    } catch (error) {
      console.error("Erreur ", error);
      return Promise.reject(error);
    }
  };
  useQuery("notifications", fetchNotifications, {
    enabled: true,
    refetchInterval: 5000,
    staleTime: 1000,
    onSuccess: (newData) => {
      setNotificationData(newData); // newData est déjà les données
    },
    onError: (error) => {
      console.error(error);
    },
  });

  return (
    <div className="w-[900px] flex flex-col space-y-2">
      {displayNotification(notificationData, sendMessage)}{" "}
      {/* Utilisez notificationData ici */}
    </div>
  );
};

export const displayNotification = (notificationData, sendMessage) => {
  return notificationData?.map((notification) => {
    return (
      <div
        key={notification.id}
        className="  flex flex-col items-start border border-gray-700 rounded  justify-start gap-2 p-1 mt-1 "
      >
        {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}
        <div className="flex">
          {/* {notification.type === "SuggestFriend"
            ? notification.from +
              " Suggests you join this group " +
              notification.type // je dois definire les valeur qui doivent afficher ici
            : ""}
          {notification.type === "JoinGroup" // je dois definire les valeur qui doivent afficher ici
            ? notification.to + " want to join  " + notification.date
            : ""}
          {notification.type === "Follow"
            ? notification.to + " want to follow you " // je dois definire les valeur qui doivent afficher ici
            : ""}
            {notification.type === "message"
            ? notification.type + " want to follow you " // je dois definire les valeur qui doivent afficher ici
            : ""} */}
            {"Notitication type: "+ notification.type + " Name sender: " + notification.first_name + " " + notification.last_name + " @"+notification.user_name}
          {/* <p className="font-semibold ">{  notification.from}</p> */}
        </div>
        <div className="flex gap-1">
          {/* <button>Going</button> */}
          <button
            className="bg-primary hover:bg-gray-600 text-sm text-white py-1 px-2 rounded flex items-center"
            onClick={() => {
              sendMessage({
                type: "ResponceNotification",
                groupeId: notification.group_id,
                id_user_sender: notification.user_id_sender,
                id_user_receiver: notification.user_id_received,
                typeNotification: notification.type,
                response: "going",
              });
              // going(notification.id, sendMessage, notification.group_id);
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
                d="M19.916 4.626a.75.75 0 0 1 .208 1.04l-9 13.5a.75.75 0 0 1-1.154.114l-6-6a.75.75 0 0 1 1.06-1.06l5.353 5.353 8.493-12.74a.75.75 0 0 1 1.04-.207Z"
                clipRule="evenodd"
              />
            </svg>
            Accept
          </button>
          <button
            onClick={() => {
              sendMessage({
                type: "ResponceNotification",
                groupeId: notification.group_id,
                id_user_sender: notification.user_id_sender,
                id_user_receiver: notification.user_id_received,
                typeNotification: notification.type,
                response: "notGoing",
              });
              // notGoing(notification.id, sendMessage, notification.group_id);type: notification.type });
            }}
            className="bg-red-400 hover:bg-gray-600 text-sm  text-white py-1 px-2 rounded flex items-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z"
                clipRule="evenodd"
              />
            </svg>
            Decline
          </button>
        </div>
      </div>
    );
    // }
  });
};

export default Notification;
