import { Metadata } from "next";
import { Poppins } from "next/font/google"
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { cookieToInitialState } from 'wagmi';
import { config } from '@/config'
import AppKitProvider from '@/contexts/WalletContext'
import { headers } from 'next/headers'

export const metadata: Metadata = {
  title: "DeVolt | The distributed charging network",
  description: "Charge your car or sell your exceeding energy in the most blockchain way possible!",
  openGraph: {
    images: "https://www.devolt.xyz/ogimage.png",
    type: "website",
    url: "https://www.devolt.xyz",
    locale: "en",
  }
};

const poppins = Poppins({
  weight: ["400", "500", "600", "700",],
  display: "swap",
  style: ["italic", "normal"],
  subsets: ["latin"],
})

const App = ({ children }: { children: React.ReactNode }) => {
  const initialState = cookieToInitialState(config, headers().get('cookie'))
  return (
    <html  lang="en" className={poppins.className}>
        <body>
          <AppKitProvider initialState={initialState}>
            {children}
            <ToastContainer theme="dark" bodyClassName={"font-medium"}/>
          </AppKitProvider>
        </body>
    </html>
  );
};

export default App;
