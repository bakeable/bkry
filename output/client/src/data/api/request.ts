import { dialogEvents } from "@/data-views/dialogs";
import { auth } from "@/firebase";
import axios from "axios";

// Get base url
const baseUrl = import.meta.env.VITE_API_ENDPOINT;

// Create config
const config = {
  baseURL: baseUrl, // url = base url + request url
  timeout: 540000,
  withCredentials: true,
  crossDomain: true,
  keepAlive: true,
};
const service = axios.create(config);

// Request interceptors
service.interceptors.request.use(
  async (config) => {
    // Add X-Access-Token header to every request, you can add other custom headers here
    const accessToken = auth.accessToken ?? useCookie("accessToken").value;
    if (accessToken) {
      config.headers.Authorization = "Bearer " + accessToken;
    }
    config.headers["Access-Control-Allow-Origin"] = window.location.origin;
    return config;
  },
  async (error: Error) => {
    console.error(error);
    await Promise.reject(error);
  },
);

// Response interceptors
service.interceptors.response.use(
  async (response) => {
    if (process.env.NODE_ENV === "development") {
      console.info(response.config.url, response);
    }
    // Get data
    const res = response.data;
    // Handle
    if (response.status >= 200 && response.status < 300) {
      return res;
    } else if (response.status === 401) {
      window.location.reload();
    } else {
      return await Promise.reject(response);
    }
  },
  async (error) => {
    console.error(error);
    if (error.code === "ERR_NETWORK") {
      document.dispatchEvent(new Event("network-error"));
    } else if (error.response.status === 417) {
      dialogEvents.ConfirmDialog.open(
        "This index does not exist yet, you will be redirected to the index creation page. Please wait 5 to 10 minutes after creating the index for the changes to take effect.",
        async () => {
          window.open(error.response.data.url, "_blank");
        },
        async () => {
          window.location.reload();
        },
      );
    } else if (error.response.status === 401) {
      document.dispatchEvent(new Event("unauthorized"));
      window.location.href =
        window.location.origin + "/login?redirect=" + window.location.pathname;
    }
    return await Promise.reject(error);
  },
);

export default service;
