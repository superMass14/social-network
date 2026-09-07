"use client";
import { useSession } from "@/app/api/api.js";
import Profile from "@/app/ui/components/profile.cp/profile.cp.js";
import { QueryClient, QueryClientProvider } from "react-query";
const queryClient = new QueryClient();
const page = () => {
  const token = localStorage.getItem("token");
  const { session, errSess } = useSession();
  const sessionId = session?.session["user_id"];
  const sessionToken = session?.session["token"];
  return (
    <QueryClientProvider client={queryClient}>
      <Profile sessionId={sessionId} sessionToken={sessionToken} />
    </QueryClientProvider>
  );
};

export default page;
