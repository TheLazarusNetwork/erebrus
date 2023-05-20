import { ConnectButton } from '@rainbow-me/rainbowkit';

const ConnectWalletButton = () => {
  return (
    <>
      <div>
        <button className="border text-blue-200 border-blue rounded-md hover:bg-black hover:border-black hover:text-black font-bold transition focus:ring focus:ring-blue-500 focus:ring-opacity-80">
          <div className="px-4 py-2">
            <ConnectButton />
          </div>
        </button>
      </div>
    </>
  );
}

export default ConnectWalletButton;
