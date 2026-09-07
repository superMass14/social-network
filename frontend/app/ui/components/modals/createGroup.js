"use client";

import { useSession } from "@/app/api/api";
import React, { useState } from "react";
import { toast } from "react-toastify";

export const CreateGroup = ({ isVisible, onClose }) => {
  const token = localStorage.getItem("token");
  const { session, errSess } = useSession();
  const [groupName, setGroupName] = useState("");
  const [groupDescription, setGroupDescription] = useState("");
  const [groupImage, setGroupImage] = useState(null);
  const [erroMsg, setErrorMsg] = useState("");

  const handleNameChange = (e) => {
    setGroupName(e.target.value);
  };

  const handleDescriptionChange = (e) => {
    setGroupDescription(e.target.value);
  };

  const handleImageChange = (e) => {
    setGroupImage(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const sessionId = session?.session["user_id"];
    // console.log(sessionId);

    const formData = new FormData();
    formData.append("name", groupName);
    formData.append("description", groupDescription);
    formData.append("image", groupImage);
    formData.append("user_id", sessionId);

    try {
      const response = await fetch("http://localhost:8080/group/creategroup", {
        method: "POST",
        body: formData,
      });
      // console.log(response.ok);
      if (response.ok) {
        setErrorMsg("");
        setGroupName("");
        setGroupDescription("");
        setGroupImage("");
        onClose();
      }
      if (response.status == 409 || response.status == 400) {
        let msg = "";

        response.status === 409
          ? (msg = "name group taken")
          : (msg = "Empty name group");

        // setErrorMsg();
        console.log(response);
        toast.error(msg, {
          position: "bottom-left",
          autoClose: 4000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
          // transition: "bounce",
        });
      }
    } catch (error) {
      console.error("Errorr:", error);
    }
  };

  if (!isVisible) return null;
  return (
    <div
      className="fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm
        flex justify-center items-center"
    >
      <div
        className="w-[700px] pb-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
             backdrop-blur-md border border-gray-700 hover:bg-opacity-95"
      >
        <button
          className="w-full p-2 flex justify-end"
          onClick={() => onClose()}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-8 h-8 hover:text-red-500
                   hover:rotate-90 transition duration-300 ease-in-out place-self-end"
          >
            <path
              fillRule="evenodd"
              d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z"
              clipRule="evenodd"
            />
          </svg>
        </button>
        <div>
          <h1 className="text-2xl text-center font-bold underline underline-offset-8 mb-5">
            Create a new group
          </h1>
          <form className="flex flex-col gap-4 px-5" onSubmit={handleSubmit}>
            <input
              type="text"
              name="name"
              required
              placeholder="Name"
              value={groupName}
              onChange={handleNameChange}
              className="bg-transparent rounded-md border border-gray-700 h-[50px]
                        focus:outline-none focus:border p-1 focus:ring-1 focus:ring-primary
                        "
            />
            <textarea
              placeholder="Group description ..."
              name="description"
              value={groupDescription}
              onChange={handleDescriptionChange}
              className="bg-transparent h-[100px] border rounded-md border-gray-700 resize-none
                        focus:outline-none focus:border p-1 focus:ring-1 focus:ring-primary
                        "
            ></textarea>
            <input
              type="file"
              name="image"
              onChange={handleImageChange}
              className="bg-transparent"
            />
            <p className="text-red-600 text-center p-2">{erroMsg}</p>
            <input
              type="submit"
              className="bg-primary rounded-md border border-gray-700 h-[50px] cursor-pointer hover:bg-second text-lg font-bold "
              value={"Create group"}
            />
          </form>
        </div>
      </div>
    </div>
  );
};
