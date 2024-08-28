import React from 'react'
import logo from "@/public/logo_horizontal.svg";
import cartesiLogo from "@/public/cartesi_icon.svg";
import Image from 'next/image';
import Link from 'next/link';

export default function Footer() {
  return (
    <footer className='bg-neutral-800 shadow-inner'>
        <div className='pt-10 flex justify-center'>
        <Image src={logo} alt='logo' height={40}/>
        </div>
        <div className='flex justify-center items-center gap-2 mt-2'>
          <p className=' text-neutral-300'>Powered by</p>
          <Link href="https://cartesi.io/" passHref>
            <Image src={cartesiLogo} alt='cartesi logo' height={30}/>
          </Link>
        </div>
        <div className='w-full flex justify-center pb-4'>
          <p className='mx-auto text-xs md:text-base text-neutral-500 p-2'>Copyright @ 2024 DevoltHQ Limited </p>
        </div>
    </footer>
  )
}
