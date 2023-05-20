import { motion } from "framer-motion";
import Header from "../components/Header";
import { Link } from "react-router-dom";
import Footer from "../components/Footer";

const LandingPage = () => {
  return (
    <>
      <Header />
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: 0.2 }}
        className="py-10"
      >
        <section className="pt-12">
          <div className="px-5 mx-auto max-w-7xl">
            <div className="w-full mx-auto text-left md:w-11/12 xl:w-9/12 md:text-center">
              <h1 className="mb-8 font-extrabold leading-none tracking-normal text-gray-100 md:tracking-tight">
                <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-blue-200 text-7xl md:text-8xl">
                  Privacy Redesigned!
                </span>
                <br />
                <span className="text-5xl md:text-6xl">
                  Secure Your Connection
                </span>
                <br />
              </h1>
              <p className="px-0 mb-8 text-xl text-gray-300 md:text-2xl lg:px-24">
                Discover Sotreus, a reliable WireGuard VPN server, simplifying
                client configurations for seamless online privacy.
              </p>
              <div className="mb-4 space-x-0 md:space-x-2 md:mb-8">
                <Link
                  to="/dashboard"
                  className="inline-flex items-center font-bold justify-center w-full px-6 py-3 mb-2 text-lg text-black rounded-2xl sm:w-auto sm:mb-0 transition bg-blue-300 hover:bg-blue-200 focus:ring focus:ring-blue-500 focus:ring-opacity-80"
                >
                  Dashboard
                </Link>
              </div>
            </div>
          </div>
        </section>
        <div className="fixed bottom-0 w-full">
          <Footer />
        </div>
      </motion.div>
    </>
  );
};

export default LandingPage;
