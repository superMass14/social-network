"use client"
import React from 'react'
import Link from 'next/link';

const Button = ({ text, onClick }) => {
    return (
        <Link href="/home">
            <button onClick={onClick} className="hover:bg-second bg-primary cursor-pointer text-text border-none w-full h-10 rounded font-bold text-center">
                {text}
            </button>
        </Link>
    );
}

export default Button