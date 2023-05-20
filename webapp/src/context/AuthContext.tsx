import React, { createContext, useContext, useState } from "react";
interface AuthContextType {
  isSignedIn: boolean;
  setIsSignedIn: React.Dispatch<React.SetStateAction<boolean>>;
  isAuthorized: boolean;
  setIsAuthorized: React.Dispatch<React.SetStateAction<boolean>>;
}

export const AuthContext = React.createContext<null | AuthContextType>(null);

type props = { children: React.ReactNode };

export const AuthProvider = ({ children }: props) => {
  const [isSignedIn, setIsSignedIn] = useState(false);
  const [isAuthorized, setIsAuthorized] = useState(false);

  return (
    <AuthContext.Provider
      value={{ isSignedIn, setIsSignedIn, isAuthorized, setIsAuthorized }}
    >
      {children}
    </AuthContext.Provider>
  );
};
