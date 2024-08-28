"use client";

import type { NextPage } from "next";
import "./globals.css";
import Topbar from "@/components/Topbar/Topbar";
import HeroSection from "@/components/HeroSection/HeroSection";
import CardsSection from "@/components/CardsSection/CardsSection";
import Footer from "@/components/Footer/Footer";
import { motion } from "framer-motion";
import SupportSection from "@/components/SupportSection/SupportSection";
import { AccordionDemo } from "@/components/Faq/Faq";
import printConsoleASCIIArt from "@/lib/ASCIIart";
import { useEffect } from "react";



const Home: NextPage = () => {

  let printedASCII = false
  useEffect(() => { !printedASCII && printConsoleASCIIArt(); printedASCII = true }, [])


  return (
    <>
      <Topbar />
      <div className="max-w-7xl mx-auto">
        <HeroSection />
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 4, duration: 1 }}
        >
          <CardsSection />
          <AccordionDemo />
          <SupportSection />
        </motion.div>
        
      </div>
      <Footer />
    </>
  );
};

export default Home;
