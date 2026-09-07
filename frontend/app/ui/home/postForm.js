"use client";
import { getSessionUser } from "@/app/_lib/utils";
import React, { use } from "react";
import { useState } from "react";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export const PostForm = (idUser, isVisible, onClose, followers) => {
  const [textarea, setTextarea] = useState(""),
    [tech, setTech] = useState(false),
    [sport, setSport] = useState(false),
    [health, setHealth] = useState(false),
    [music, setMusic] = useState(false),
    [news, setNews] = useState(false),
    [other, setOther] = useState(true),
    [privacy, setPrivacy] = useState();
  //  [imgs, setImgs] = useState(null),

  const handlePost = async (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    data.append("userID", idUser);
    data.append("type", "createPost");
    // console.log("my data => ", data);
    const options = {
      method: "POST",
      body: data,
    };
    fetch("http://localhost:8080/post", options).then(async (x) => {
      const retrieved = await x.json();
      // console.log("response", retrieved);
      if (retrieved.type != "success") {
        toast.error(retrieved.msg, {
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
        return;
      }else{
        onClose()
      }
      //!emptying inputs after submit
      isVisible = true;
      setTextarea("");
      setSport(false);
      setTech(false);
      setSport(false);
      setHealth(false);
      setMusic(false);
      setNews(false);
      setOther(true);
    });
  };

  if (!isVisible) return null;

  return (
    <div
      className="fixed inset-0 z-30 bg-bg bg-opacity-10 backdrop-blur-sm
        flex justify-center items-center"

      // onClick={() => onClose()}
    >
      <div
        className="w-[700px] h-[300px] p-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
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
        <div className="flex flex-col items-center">
          <h1 className="text-2xl text-center font-bold underline underline-offset-8 mb-5">
            Create a new post
          </h1>
          <form
            className="flex flex-col bg-bg gap-1  "
            method=""
            data-form="post"
            encType="multipart/form-data"
            onSubmit={handlePost}
          >
            <TextArea
              label="Post Title"
              name="content"
              placeholder="Let's post something"
              required
              value={textarea}
              onChange={(e) => setTextarea(e.target.value)}
            />
            <div className="flex items-start justify-end">
              <div className="flex gap-1 flex-wrap mr-2 mt-1 text-sm">
                <Checkbox
                  label="Tech"
                  value="Tech"
                  name="Tech"
                  checked={tech}
                  onChange={() => (tech ? setTech(false) : setTech(true))}
                />
                <Checkbox
                  label="Sport"
                  value="Sport"
                  name="Sport"
                  checked={sport}
                  onChange={() => (sport ? setSport(false) : setSport(true))}
                />
                <Checkbox
                  label="Health"
                  value="Health"
                  name="Health"
                  checked={health}
                  onChange={() => (health ? setHealth(false) : setHealth(true))}
                />
                <Checkbox
                  label="Music"
                  value="Music"
                  name="Music"
                  checked={music}
                  onChange={() => (music ? setMusic(false) : setMusic(true))}
                />
                <Checkbox
                  label="News"
                  value="News"
                  name="News"
                  checked={news}
                  onChange={() => (news ? setNews(false) : setNews(true))}
                />
                <Checkbox
                  label="Other"
                  value="Other"
                  name="Others"
                  checked={other}
                  onChange={() => (other ? setOther(false) : setOther(true))}
                />
              </div>

              <PrivacySelect followers={followers} />

              <label htmlFor="image_post" className=" cursor-pointer mr-2">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  className="w-8 h-8"
                >
                  <path
                    fillRule="evenodd"
                    d="M1.5 6a2.25 2.25 0 0 1 2.25-2.25h16.5A2.25 2.25 0 0 1 22.5 6v12a2.25 2.25 0 0 1-2.25 2.25H3.75A2.25 2.25 0 0 1 1.5 18V6ZM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0 0 21 18v-1.94l-2.69-2.689a1.5 1.5 0 0 0-2.12 0l-.88.879.97.97a.75.75 0 1 1-1.06 1.06l-5.16-5.159a1.5 1.5 0 0 0-2.12 0L3 16.061Zm10.125-7.81a1.125 1.125 0 1 1 2.25 0 1.125 1.125 0 0 1-2.25 0Z"
                    clipRule="evenodd"
                  />
                </svg>
              </label>
              <input type="file" name="image" id="image_post" />
              <button
                type="submit"
                className="bg-second h-10 text-lg font-bold pl-3 pr-3 rounded-lg cursor-pointer hover:bg-primary"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  className="w-6 h-6"
                >
                  <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
                </svg>
              </button>
              {/* < Button text="Log In" onClick={handlePost()} /> */}
            </div>{" "}
          </form>
        </div>
      </div>
    </div>
  );
};

export const PrivacySelect = ({ followers }) => {
  const [selectedValue, setSelectedValue] = useState("public");
  const [showUserList, setShowUserList] = useState(false);
  const [selectedUsers, setSelectedUsers] = useState([]);

  const handleChange = (event) => {
    setSelectedValue(event.target.value);
    setShowUserList(event.target.value === "almost");
  };

  const handleUserSelection = (userId) => {
    const isSelected = selectedUsers.includes(userId);
    setSelectedUsers(
      isSelected
        ? selectedUsers.filter((id) => id !== userId)
        : [...selectedUsers, userId]
    ); // Toggle user selection
  };

  return (
    <div className="flex flex-col items-center mr-2">
      <select
        id="privacy-select"
        value={selectedValue}
        name="privacy"
        onChange={handleChange}
        className="w-32 rounded-md px-2 py-1 font-bold outline-none focus:ring-1 bg-primary focus:ring-primary"
      >
        <option value="public">Public</option>
        <option value="private">Private</option>
        <option value="almost">Select users</option>
      </select>
      <input type="hidden" value={selectedUsers} name="viewers" />
      {showUserList && (
        <div className="mt-1 max-h-44  w-[300px] p-2 overflow-scroll border rounded-md">
          <ul className="flex flex-wrap gap-2">
            {followers?.map((user) => (
              <li key={user.id}>
                <label className="flex gap-1 cursor-pointer">
                  <input
                    type="checkbox"
                    name={`view-${user.first_name}`}
                    //checked={selectedUsers.includes(user.name)}
                    value={user.id}
                    onChange={() => handleUserSelection(user.id)}
                  />
                  {user.first_name + " " + user.last_name}
                </label>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

// const followers = [
//   { name: "Vindour", src: "/assets/profilibg.jpg", alt: "profil", id: 3 },
//   { name: "ibg", src: "/assets/profilibg.jpg", alt: "profil", id: 2 },
//   { name: "nixa", src: "/assets/profilibg.jpg", alt: "profil", id: 4 },
//   { name: "dickss", src: "/assets/profilib g.jpg", alt: "profil", id: 1 },
// ];

export function TextArea({
  label,
  name,
  placeholder,
  required,
  value,
  onChange,
}) {
  return (
    <div className="mb-2">
      {/* {<label className="block text-sm font-medium text-gray-700 mb-1" htmlFor={name}>{label}</label>} */}
      <textarea
        className="resize-none w-[100%] h-10 border border-gray-700 focus:outline-none focus:border bg-transparent text-text rounded-md p-1 focus:ring-1 focus:border-primary focus:ring-primary"
        name={name}
        placeholder={placeholder}
        required={required}
        value={value}
        onChange={onChange}
      />
    </div>
  );
}

export function Checkbox({ label, value, name, checked, onChange }) {
  return (
    <div className="checkbox-container">
      <input
        type="checkbox"
        id={label}
        value={value}
        name={name}
        checked={checked}
        onChange={onChange}
      />
      <label htmlFor={label}>{label}</label>
    </div>
  );
}
