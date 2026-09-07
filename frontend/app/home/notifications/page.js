"use client";
// import { ApiProvider } from "@/app/_lib/utils";
import Notification from "@/app/ui/home/notification/page";
import React from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "../groups/page";

const Page = () => {
  return (
    <>
      {/* <ApiProvider> */}
        <QueryClientProvider client={queryClient}>
          <Notification />
        </QueryClientProvider>
      {/* </ApiProvider> */}
    </>
  );
};

export default Page;
