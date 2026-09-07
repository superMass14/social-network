import Link from "next/link";
import { FaSearch } from "react-icons/fa";
import { AiFillMessage } from "react-icons/ai";
import { IoNotifications } from "react-icons/io5";

export function Header() {
  return (
    <div className="shadow  w-screen z-10 h-16 p-4 flex justify-between items-center">
      <Link href="/home" className="">
        <img
          src="/assets/logo.png"
          alt="logo"
          width={100}
          height={200}
          className=""
        />
      </Link>
      <div className="flex xl:w-[650px] sm:w-96 h-10 w-40 border border-gray-700 items-center gap-x-2  color-black p-2 rounded-md">
        <i className="">
          <FaSearch className=" text-2xl" />
        </i>
        <input
          type="search"
          placeholder="Search ..."
          name="search"
          className="h-8 w-full bg-transparent text-md focus:outline-none"
        />
      </div>
      <div className="flex items-center text-3xl gap-x-5 mr-3">
        <i className="hover:text-second cursor-pointer">
          <AiFillMessage />
        </i>
        <i className="hover:text-second cursor-pointer ">
          <IoNotifications />
        </i>
      </div>
    </div>
  );
}
