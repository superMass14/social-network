"use client";
import { useEffect } from "react";
import { useState } from "react";
import { Header } from "../ui/components/header/header";
import Navbar from "../ui/components/navbar/navbar";
import Sidebar from "../ui/components/sidebarRight/sidebar";
import { displayFollowers } from "./utils";
import { ToastContainer } from "react-toastify";
import { useSession } from "../api/api";

export default function Layout({ children }) {

  return (
    <div className="h-screen">
      <div className="fixed">
        <Header />
      </div>
      <div
        className="md:flex md:flex-row flex flex-col-reverse 
      justify-between h-[99%] md:justify-between md:h-full  overflow-hidden"
      >
        <div className="md:mt-20">
          <Navbar/>
          <ToastContainer />
        </div>{" "}
        {/* Enveloppez le contenu avec WebSocketProvider */}
        <div className="mt-20 overflow-x-hidden overflow-y-scroll chrome pl-3 mr-3">
          {children}
        </div>
        <div className="md:mt-20 hidden md:block">
          <Sidebar
          />
        </div>
      </div>
    </div>
  );
}
