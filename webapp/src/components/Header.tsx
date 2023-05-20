import { useNavigate } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../context/AuthContext";
import { useAccount, useSignMessage, useNetwork } from "wagmi";
import { getChallengeId, getToken } from "../modules/api";
import { ConnectButton } from "@rainbow-me/rainbowkit";

const Header = () => {
  const navigate = useNavigate();
  const authContext = useContext(AuthContext);
  const useraddress = useAccount();
  const [message, setMessage] = useState<string>("");
  const [challengeId, setChallengeId] = useState<string>("");
  const [signature, setSignature] = useState<string | undefined>();
  const { signMessageAsync } = useSignMessage();
  const { isConnected, address } = useAccount();

  const navigateDashboard = async () => {
    await navigate("/dashboard");
  };

  const navigateServer = async () => {
    await navigate("/server");
  };

  useEffect(() => {
    const storedToken = localStorage.getItem("token");
    if (storedToken) {
      authContext?.setIsSignedIn(true);
    }
    let timeoutId: string | number | NodeJS.Timeout | null = null;

    const getSignMessage = async () => {
      if (address == undefined) {
        if (timeoutId !== null) {
          clearTimeout(timeoutId);
        }

        timeoutId = setTimeout(() => {
          signOut();
        }, 500);
      } else if (!authContext?.isSignedIn) {
        if (timeoutId !== null) {
          clearTimeout(timeoutId);
        }

        const response = await getChallengeId(address);
        setMessage(response.data.eula + response.data.challangeId);
        setChallengeId(response.data.challangeId);
        if (response.data.isAuthorized == true) {
          authContext?.setIsAuthorized(true);
        } else {
          authContext?.setIsAuthorized(false);
        }
      }
    };

    getSignMessage();

    return () => {
      if (timeoutId !== null) {
        clearTimeout(timeoutId);
      }
    };
  }, [authContext?.isSignedIn, address]);

  const signMessage = async () => {
    const signature = await signMessageAsync({ message });
    setSignature(signature);
    //make a post request to the sotreus server with the signature and challengeId

    const response = await getToken(signature, challengeId);
    if (response.data.token) {
      //store the token in the session storage
      sessionStorage.setItem("token", response.data.token);
      localStorage.setItem("token", response.data.token);
      authContext?.setIsSignedIn(true);
    }
  };
  const signOut = () => {
    sessionStorage.removeItem("token");
    localStorage.removeItem("token");
    setMessage("");
    setSignature("");
    setChallengeId("");
    authContext?.setIsSignedIn(false);
  };

  return (
    <div>
      <div className="navbar">
        <div className="flex-1">
          <a
            href="/"
            className="btn btn-ghost normal-case text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-blue-200 text-2xl md:text-3xl"
          >
            Sotreus
          </a>
        </div>
        <div className="flex-none">
          <ul className="menu menu-horizontal px-1">
            <div className="flex space-x-2">
              {(!address || authContext?.isSignedIn) && (
                <div className="border text-blue-200 border-blue rounded-md hover:bg-black hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80">
                  <div className="px-4 py-2">
                    <ConnectButton />
                  </div>
                </div>
              )}
              {!(isConnected && authContext?.isSignedIn) && address && (
                <li>
                  <button
                    onClick={signMessage}
                    className="border text-blue-200 border-blue hover:bg-blue-300 hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80"
                  >
                    Sign In
                  </button>
                </li>
              )}
              <li>
                <button
                  onClick={navigateDashboard}
                  className="border text-blue-200 border-blue hover:bg-blue-300 hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80"
                >
                  Dashboard
                </button>
              </li>{" "}
            </div>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Header;
