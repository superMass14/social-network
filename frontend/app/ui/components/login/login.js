"use client";
import { loginUser, logoutUser } from "@/app/api/api.js";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Login() {
  const [errorMessage, setErrorMessage] = useState(null);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await loginUser(email, password);
      if (response.error === null && response.data) {
        localStorage.setItem("token", response.data.token);
        router.push("/home");
      } else {
        setErrorMessage(response.error);
      }
    } catch (error) {
      console.error(error);
    }
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
            Don't have an account ?
            <Link
              href="/register"
              className="text-primary hover:text-second cursor-pointer"
            >
              Sign Up.
            </Link>
          </div>
        </div>
        <div className="mt-40 sm:mt-56 flex flex-col gap-3 w-8/12 max-w-96">
          <h1 className="text-center font-bold text-4xl">Welcome Back</h1>
          <div className="text-center login_other">
            <h4 className="font-bold text-xl mb-2">Login in to account</h4>
            <div className="flex items-center justify-center gap-4">
              <img
                src="/assets/login/google.svg"
                alt="google"
                width={40}
                height={40}
              />
              <img
                src="/assets/login/githubb.svg"
                alt="google"
                width={40}
                height={40}
              />
            </div>
          </div>
          <form
            onSubmit={handleLogin}
            className="form_login flex flex-col gap-3"
          >
            <input
              type="email"
              name="email"
              placeholder="Email"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />

            <input
              type="password"
              name="password"
              placeholder="Password"
              required
              className="h-10 rounded pl-2 border border-border_color text-bg"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
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
              type="submit"
              className="hover:bg-second bg-primary cursor-pointer text-text border-none w-full h-10 rounded font-bold text-center"
            >
              Log In
            </button>
          </form>
        </div>
      </div>
      <div className="bg-[url('/assets/login/bg.jpg')] bg-cover bg-center w-6/12 h-screen hidden sm:block"></div>
    </div>
  );
}
