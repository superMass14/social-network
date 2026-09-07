"use client";
import { registerUser } from "@/app/api/api.js";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

function Register() {
  const [errorMessage, setErrorMessage] = useState(null);
  const [pending, setPending] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setPending(true);
    try {
      let data = new FormData(e.target);
      const response = await registerUser(data);
      console.log("res => ", response)
      if (response.error == "ok") {
        router.push("/login");
      } else {
        setErrorMessage(response.error);
      }
    } catch (error) {
      console.error("Something went wrong. Please try again.", error);
    }

    setPending(false);
  };

  return (
    <div className="w-screen h-screen flex">
      <div className="flex flex-col items-center w-full sm:w-6/12">
        <div className="flex items-center justify-around w-full">
          <img
            src="/assets/icons/comment.png"
            alt="logo"
            width={40}
            height={40}
          />
          <div>
            You already have an account?{" "}
            <Link
              href="/login"
              className="text-primary hover:text-second cursor-pointer"
            >
              Sign In.
            </Link>
          </div>
        </div>
        <div className="mt-2 sm:mt-24 flex flex-col gap-3 w-8/12 max-w-96">
          <h1 className="text-center font-bold text-4xl">Sign Up</h1>
          <div className="flex items-center justify-center gap-4">
            <img src="/assets/login/google.svg" alt="google" />
            <img src="/assets/login/githubb.svg" alt="github" />
          </div>
          <form className="w-full flex flex-col gap-3" onSubmit={handleSubmit}>
            <input
              type="text"
              name="firstname"
              placeholder="Firstname"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <input
              type="text"
              name="lastname"
              placeholder="Lastname"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <input
              type="date"
              name="dateofbirth"
              placeholder="Date of birth"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <input
              type="file"
              name="avatar"
              placeholder="Avatar"
              className="rounded border border-border_color text-bg"
            />
            <input
              type="text"
              name="username"
              placeholder="Username"
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <input
              type="email"
              name="email"
              placeholder="E-mail"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <input
              type="password"
              name="password"
              placeholder="Password"
              required
              className="h-10 w-full rounded pl-2 border border-border_color text-bg"
            />
            <textarea
              type="text"
              name="aboutme"
              placeholder="About me"
              maxLength={1500}
              className="h-20 pt-6 rounded pl-2 border border-border_color text-bg resize-none"
            />

            <div>
              {errorMessage && (
                <div className="items-center w-full bg-red-100 border border-red-400 rounded-md py-2 px-3 mb-4 text-red-700">
                  <strong className="font-bold">Wrong Credentials </strong>
                  <br />
                  <span className="block sm:inline">{errorMessage}</span>
                </div>
              )}
            </div>
            <button
              className="hover:bg-second bg-primary cursor-pointer border-none w-full h-10 rounded font-bold text-text text-center"
              disabled={pending}
            >
              Create account
            </button>
          </form>
        </div>
      </div>
      <div className="bg-[url('/assets/login/bg.jpg')] bg-cover bg-center w-6/12 h-screen hidden sm:block"></div>
    </div>
  );
}

export default Register;
