import React from "react";

export const DisplayMembers = ({ members, isVisible, onClose }) => {
  if (!isVisible) return null;
  return (
    <div
      className="fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm 
        flex justify-center items-center"
    >
      <div
        className="w-[700px] h-[600px] pb-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
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
            Groups members
          </h1>
          <div className="flex flex-col lg:w-[100%] 2xl-[80%] xl:w-[75%] w-[80%]  gap-1 ">
            <input
              placeholder="search friend ..."
              className="h-8 bg-transparent border border-gray-700 rounded-md text-center focus:outline-none focus:border-primary"
            />
            <div className="flex flex-col h-[400px] overflow-scroll">
              {members ? displaySuggestFriend(members) : ""}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export const displaySuggestFriend = (data) => {
  return data.map((follower) => {
    return (
      <div
        key={follower.email}
        className=" hover:opacity-90 w-[95%] flex items-center cursor-pointer justify-between gap-2 mt-1 mb-3 p-2  "
      >
        {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}
        <div className="flex items-center gap-2">
          <img
            className="rounded-full "
            src={
              follower.avatar !== "NULL"
                ? follower.avatar
                : "/assets/profilibg.jpg"
            }
            alt={follower.avatar}
            width={40}
            height={40}
          />
          <h4 className="font-bold ">{follower.first_name}</h4>
        </div>
      </div>
    );
  });
};
