import Header from "./Header";

const NotAuthorized = () => {
  return (
    <div className="flex h-screen flex-col">
      <Header />
      <div className="flex flex-col items-center justify-center flex-1">
        <div className="text-4xl md:text-6xl font-bold text-blue-200 mb-4">
          Not Authorized
        </div>
        <div className="text-xl md:text-2xl text-white">
          Your wallet address is not authorized
        </div>
      </div>
    </div>
  );
};

export default NotAuthorized;
