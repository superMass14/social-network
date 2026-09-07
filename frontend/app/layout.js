'use client'
import "./globals.css";
import { WebSocketProvider } from "./_lib/websocket"; // Importez WebSocketProvider

// export const metadata = {
//  title: "Social-Network",
//  description: "Welcome to SNK, where every word is a work of art",
// };

export default function RootLayout({ children }) {
 return (
    <html lang="en">
      {/* <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/icon?family=Material+Icons"
      /> */}
      <WebSocketProvider> {/* Enveloppez le contenu avec WebSocketProvider */}
        <body className="">{children}</body>
      </WebSocketProvider>
    </html>
 );
}