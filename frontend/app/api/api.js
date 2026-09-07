"use client";
import { useEffect, useState } from "react";
const NEXT_PUBLIC_API_URL = `http://localhost:8080`;

export async function loginUser(email, password) {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signin`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  })
    .then((response) => {
      if (response.status === 401) {
        return { error: "Invalid email or password." };
      }
      if (response.ok) {
        return response.json().then((json) => {
          localStorage.setItem("token", json.token);
          return { error: null, data: json };
        });
      }
      throw new Error("An error occurred.");
    })
    .catch((error) => {
      console.error(error);
      return { error: "An error occurred. Please try again." };
    });
}
export async function logoutUser(token) {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signout`, {
    method: "DELETE",
    headers: {
      Authorization: token,
    },
  })
    .then((response) => response.json())
    .catch((error) => {
      return { error: "Not authorized!!!" };
    });
}

export async function registerUser(formData) {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signup`, {
    method: "POST",
    body: formData,
  })
    .then(async (response) => {
      const json = await response.json();
      if (response.ok) {
        return { error: "ok", data: json };
      }
      return { error: json.error };
    })
    .catch((error) => {
      return { error: error };
    });
}
export function useSession() {
  const [session, setSession] = useState(null);
  const [error, setError] = useState(null);

  const token = localStorage.getItem('token')
  useEffect(() => {
    fetchSessionData();
  }, [token]);

  async function fetchSessionData() {
    try {
      const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
        method: "GET",
        headers: {
          Authorization: token,
        },
      });
      if (!response.ok) {
        throw new Error(
          `Failed to fetch session data. Status: ${response.status}`
        );
      }

      const sessionData = await response.json();
      setSession(sessionData);
    } catch (error) {
      console.error(error);
      setError(error);
    }
  }

  return { session, error };
}

export async function getProfileById(id, sessionId, token) {
  try {
    const response = await fetch(
      `${NEXT_PUBLIC_API_URL}/profile/user?id=${id}&sessionId=${sessionId}`,
      {
        method: "GET",
        headers: {
          Authorization: token,
        },
      }
    );
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function followUser(id, sessionId, token) {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/follow?id=${id}`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({
        followed_id: String(id),
        follower_id: String(sessionId),
      }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function unfollowUser(id, sessionId, token) {
  // alert("unfollow");
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/unfollow?id=${id}`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({
        followed_id: String(id),
        follower_id: String(sessionId),
      }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function profileTypeUser(sessionId, user_type, token) {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/profile/type`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({ sessionId: String(sessionId), user_type }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}
