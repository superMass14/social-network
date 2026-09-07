'use client'

import React, { useState } from "react";
import { toast } from "react-toastify";

export const CreateEvent = ({ isVisible, onClose, id, user }) => {
  const [textarea, setTextarea] = useState("");
  const [date, setDate] = useState("");

  const [selectedDate, setSelectedDate] = useState("");
  const [selectedHour, setSelectedHour] = useState("");

  const handleHourChange = (event) => {
    setSelectedHour(event.target.value);
  };

  // Function to combine date and time for internal use (optional)
  const getCombinedDateTime = () => {
    if (selectedDate && selectedHour) {
      return `${selectedDate}T${selectedHour}:00`;
    }
    return ""; // Return empty string if both aren't selected
  };

  const handleDateChange = (event) => {
    const selectedDate = new Date(event.target.value);
    const today = new Date();

    if (selectedDate > today) {
      setDate(event.target.value);
    }
  };

  const handlePost = (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    const token = localStorage.getItem("token") || null;
    const options = {
      method: "POST",
      body: data,
    };
    fetch(
      `http://localhost:8080/group/createEvent?id=${id}&token=${token}`,
      options
    ).then(async (x) => {
      const retrieved = await x.json();
      onClose();

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
      }
      setTextarea("");
      setDate("");
    });
  };

  if (!isVisible) return null;
  return (
    <div
      className="fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm
        flex justify-center items-center"

      // onClick={() => onClose()}
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
            Create an Event
          </h1>
          <form
            className="flex flex-col gap-4 px-5"
            onSubmit={handlePost}
            data-form="post"
            encType="multipart/form-data"
          >
            <textarea
              placeholder="Event description ..."
              required
              className="bg-transparent h-[100px] border rounded-md border-gray-700 resize-none
                        focus:outline-none  p-1 focus:ring-1 focus:ring-primary
                        "
              name="content"
              value={textarea}
              onChange={(e) => setTextarea(e.target.value)}
            ></textarea>
            <label htmlFor="date" className="text-gray-300">
              Choose a date
            </label>
            <input
              type="date"
              name="date"
              value={date}
              required
              onChange={handleDateChange}
              className="border border-gray-700 bg-transparent focus:outline-none p-1 focus:ring-1 focus:ring-primary mr-2" // Add margin for spacing
            />
            <select
              required
              name="heure"
              className="border border-gray-700 bg-transparent focus:outline-none p-1 focus:ring-1 focus:ring-primary mr-2"
              value={selectedHour}
              onChange={handleHourChange}
            >
              {Array.from({ length: 24 }, (_, i) => ({
                value: i < 10 ? `0${i}` : `${i}`,
                label: `${i}:00`,
              })).map((hour) => (
                <option key={hour.value} value={hour.value}>
                  {hour.label}
                </option>
              ))}
            </select>
            {/* <input type='file' className='bg-transparent' /> */}
            <input
              type="submit"
              className="bg-primary rounded-md border border-gray-700 h-[50px] cursor-pointer hover:bg-second text-lg font-bold "
              value={"Create Event"}
            />
          </form>
        </div>
      </div>
    </div>
  );
};
