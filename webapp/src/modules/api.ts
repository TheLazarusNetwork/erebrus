import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { getBaseUrl } from './Utils';

const baseURL = getBaseUrl()

export interface UpdateClientPayload {
  id: string;
  name: string;
  type: string;
  email: string;
  enable: boolean;
  ignorePersistentKeepalive: boolean;
  presharedKey: string;
  allowedIPs: string[];
  address: string[];
  privateKey: string;
  publicKey: string;
  createdBy: string;
  updatedBy: string;
  created: string;
  updated: string;
}

export interface CreateClientPayload {
  name: string;
  tags: string[];
  email: string;
  enable: boolean;
  allowedIPs: string[];
  address: string[];
  createdBy: string;
}
interface ClientResponse {
  status: number;
  success?: boolean; 
  sucess?: boolean; 
  message: string;
  clients: any[];
}

interface ChallengeidResponse {
  status: number;
}

// ========================================( DASHBOARD APIs )==================================================== //

async function callSotreusAPI(
  endpoint: string,
  method: 'GET' | 'POST' | 'DELETE' | 'PATCH',
  payload?: CreateClientPayload | UpdateClientPayload | null,
  clientId?: string,
  qrcode?: boolean
): Promise<AxiosResponse<any>> {
  const axiosInstance = axios.create({
    baseURL,
  });

  const config: AxiosRequestConfig = {
    method,
    url: clientId ? endpoint.replace(':client_id', clientId) : endpoint,
    headers: {
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }
  };

  if (qrcode) {
    config.params = { qrcode: true };
  }

  if (payload) {
    config.data = payload;
  }

  return axiosInstance(config);
}

export async function emailClientConfig(clientId: string): Promise<AxiosResponse<any>> {
  return axios.get(`${baseURL}/api/v1.0/client/${clientId}/email`,
  {headers:{
    "Authorization": `Bearer ${localStorage.getItem("token")}`
  }})
}

export const updateServer = async (updatedConfig: any) => {
  const response = await axios.patch(`${baseURL}/api/v1.0/server`, updatedConfig,
  {headers:{
    "Authorization": `Bearer ${localStorage.getItem("token")}`
  }});
  return response.data;
};

export async function getClientInfo(clientId: string): Promise<AxiosResponse<any>> {
  return callSotreusAPI('/api/v1.0/client/:client_id', 'GET', null, clientId);
}

export async function createClient(payload: CreateClientPayload): Promise<AxiosResponse<any>> {
  return callSotreusAPI('/api/v1.0/client', 'POST', payload, );
}

export async function updateClient(clientId: string, payload: UpdateClientPayload): Promise<AxiosResponse<any>> {
  return axios.patch(`${baseURL}/api/v1.0/client/${clientId}`,payload, {headers:{
    "Authorization": `Bearer ${localStorage.getItem("token")}`
  }})
}

export async function getClients(token: string | null): Promise<ClientResponse> { 
  const url = `${baseURL}/api/v1.0/client`
  const response = await axios.get<ClientResponse>(url, {headers:{
    "Authorization": `Bearer ${localStorage.getItem("token")}`
  }});
  if (response.status === 200) {
    return response.data;
  } else {
    throw new Error(`Request failed with status: ${response.status}`);
  }
}

export async function deleteClient(clientId: string): Promise<AxiosResponse<any>> {
  return callSotreusAPI('/api/v1.0/client/:client_id', 'DELETE', null, clientId);
}

export async function getClientConfig(clientId: string, qrcode?: boolean): Promise<AxiosResponse<any>> {
  return callSotreusAPI('/api/v1.0/client/:client_id/config', 'GET', null, clientId, qrcode);
}

// =============================================( SERVER APIs )=================================================== //

export const getStatus = async () => {
    const response = await axios.get(`${baseURL}/api/v1.0/status`);
    return response.data;
};

export const getServerInfo = async () => {
  const response = await axios.get(`${baseURL}/api/v1.0/server`, {headers:{
    "Authorization": `Bearer ${localStorage.getItem("token")}`
  }});
  return response.data;
};

export const getServerConfig = async () => {
    const response = await axios.get(`${baseURL}/api/v1.0/server/config`, {headers:{
      "Authorization": `Bearer ${localStorage.getItem("token")}`
    }});
    return response.data;
};
    
export const getChallengeId = async (address: `0x${string}` | undefined) => {
  let response;
  try {
    // Make a get request to your server
    response = await axios.get(
      `${baseURL}/api/v1.0/authenticate?walletAddress=${address}`,
      {
        headers: {
          "Content-Type": "application/json",
        },  
      }
    );
  } catch (error) {
    throw error;
  }
  return response;
};

export const getToken = async (signature: string | undefined, challengeId:string) => {
  let response;
  try {
    // Make a post request to your server
    response = await axios.post(
      `${baseURL}/api/v1.0/authenticate`,
      { signature,challengeId },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
  } catch (error) {
    throw error;
  }
  return response;
};