"use client";
import { useSession } from "@/app/api/api.js";
import Link from "next/link.js";
import { useRouter } from "next/navigation.js";
import { useEffect } from "react";

export function Landing() {
  const router = useRouter();
  const token = localStorage.getItem("token") || null;
  const { session, error } = useSession();
  useEffect(() => {
    if (session !== null) {
      router.push("/home");
    }
  }, [session]);
  return (
    <div className="">
      <div className=" h-screen">
        <header className="">
          <div className="px-4 mx-auto sm:px-6 lg:px-8">
            <div className="flex items-center justify-between h-16 lg:h-20">
              <div className="flex-shrink-0">
                <a href="#" title="" className="flex">
                  {/* <img
                    className="w-auto h-8"
                    src="https://cdn.rareblocks.xyz/collection/celebration/images/logo.svg"
                    alt=""
                  /> */}
                </a>
              </div>
              <button
                type="button"
                className="inline-flex p-2 text-black transition-all duration-200 rounded-md lg:hidden focus:bg-gray-100 hover:bg-gray-100"
              >
                {/* Menu open: "hidden", Menu closed: "block" */}
                <svg
                  className="block w-6 h-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M4 8h16M4 16h16"
                  />
                </svg>
                {/* Menu open: "block", Menu closed: "hidden" */}
                <svg
                  className="hidden w-6 h-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M6 18L18 6M6 6l12 12"
                  />
                </svg>
              </button>
              <Link
                href="/register"
                title=""
                className="hidden lg:inline-flex items-center justify-center px-5 py-2.5 text-base transition-all duration-200 hover:bg-second focus:text-black font-semibold text-white bg-primary rounded-full"
                role="button"
              >
                {" "}
                Join Now{" "}
              </Link>
            </div>
          </div>
        </header>
        <section className="bg-opacity-30 py-10 sm:py-16 lg:py-24">
          <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
            <div className="grid items-center grid-cols-1 gap-12 lg:grid-cols-2">
              <div>
                <p className="text-base font-semibold tracking-wider text-second uppercase">
                  Welcome to SNK
                </p>
                <h1 className="mt-4 text-4xl font-bold lg:mt-8 sm:text-6xl xl:text-7xl">
                  Share your passions, inspire the world
                </h1>
                <Link
                  href="/register"
                  title=""
                  className="inline-flex items-center px-6 py-4 mt-8 font-semibold transition-all duration-200 text-text bg-primary rounded-full lg:mt-16 hover:bg-second"
                  role="button"
                >
                  Join for free
                  <svg
                    className="w-6 h-6 ml-8 -mr-2"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="1.5"
                      d="M13 9l3 3m0 0l-3 3m3-3H8m13 0a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                </Link>
                <p className="mt-5 text-">
                  Already joined us?{" "}
                  <Link
                    href="/login"
                    title=""
                    className="text-second transition-all duration-200 hover:underline hover:text-primary"
                  >
                    Log in
                  </Link>
                </p>
              </div>
              <div>
                <img
                  className="w-full"
                  src="https://cdn.rareblocks.xyz/collection/celebration/images/hero/1/hero-img.png"
                  alt=""
                />
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>
  );
}
