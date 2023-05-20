import Header from "./Header";

const NotSigned = ({ component }: { component: string }) => {
  return (
    <div className="flex h-screen flex-col">
      <Header />
      <div className="flex flex-col items-center justify-center flex-1">
        <div className="text-4xl md:text-6xl font-bold text-blue-200 mb-4">
          Not Signed in
        </div>
        <div className="text-xl md:text-2xl text-white">
          Please sign in to access the {component}
        </div>
      </div>
    </div>
  );
};

export default NotSigned;
