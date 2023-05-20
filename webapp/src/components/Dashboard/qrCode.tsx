import React, { useEffect, useState } from "react";
import QRCode from "qrcode.react";
import { getClientConfig } from "../../modules/api";

interface QrCodeProps {
  clientId: string;
}

export const QrCode: React.FC<QrCodeProps> = ({ clientId }) => {
  const [qrCodeData, setQrCodeData] = useState(null);

  const handleQrCode = async (clientId: string) => {
    try {
      const response = await getClientConfig(clientId, false);
      setQrCodeData(response.data);
    } catch (error) {}
  };

  useEffect(() => {
    handleQrCode(clientId);
  }, [clientId]);

  return (
    <div>
      <label htmlFor="my-modal-3">Qr Code</label>
      <input type="checkbox" id="my-modal-3" className="modal-toggle" />
      <div className="modal ">
        <div className="modal-box relative w-full bg-gray-900 rounded-md bg-clip-padding backdrop-filter backdrop-blur-sm bg-opacity-80 border border-gray-100">
          <label
            htmlFor="my-modal-3"
            className="btn btn-sm btn-circle absolute right-2 top-2"
          >
            ✕
          </label>
          <h3 className="text-3xl text-blue-200 font-bold">
            Scan the following Qr Code!
          </h3>
          <p className="py-4 text-sm text-white font-medium">
            You will get the client's config qrcode here..
          </p>
          <div className="mt-5 flex justify-center">
            {qrCodeData && <QRCode value={qrCodeData} size={450} />}
          </div>
        </div>
      </div>
    </div>
  );
};

export default QrCode;
