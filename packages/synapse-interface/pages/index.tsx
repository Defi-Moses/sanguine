import Head from 'next/head'
import { useAccount } from 'wagmi'

import BridgePage from './bridge'

const Home = () => {
  const { address } = useAccount()
  return (
    <>
      <Head>
        <title>Synapse</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BridgePage address={address} />
    </>
  )
}

export default Home