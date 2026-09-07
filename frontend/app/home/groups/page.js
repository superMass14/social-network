'use client'
// import { WebSocketProvider } from '@/app/_lib/websocket'
import React from 'react'

import { QueryClient } from "react-query";
import { QueryClientProvider } from "react-query";
// import { ApiProvider } from "@/app/_lib/utils";
import Groups from '@/app/ui/home/groups/groups';

export const queryClient = new QueryClient();

const page = () => {
  return (
    <>
      {/* <ApiProvider> */}
        <QueryClientProvider client={queryClient}>
          <Groups />
        </QueryClientProvider>
      {/* </ApiProvider> */}
    </>
  );
};

export default page;
