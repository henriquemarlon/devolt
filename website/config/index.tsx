import { defaultWagmiConfig } from '@web3modal/wagmi/react/config'
import { cookieStorage, createStorage } from 'wagmi'
import { arbitrumSepolia } from 'wagmi/chains'

export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID_KEY; 

if (!projectId) throw new Error('Project ID is not defined')

export const metadata = {
  name: 'AppKit',
  description: 'AppKit Example',
  url: 'https://web3modal.com', 
  icons: ['https://avatars.githubusercontent.com/u/37784886']
}

const chains = [arbitrumSepolia] as const
export const config = defaultWagmiConfig({
  chains,
  projectId,
  metadata,
  ssr: true,
  storage: createStorage({
    storage: cookieStorage
  }),
})