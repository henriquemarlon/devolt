"use client";
import React, { useState } from "react";
import type { NextPage } from "next";
import "../globals.css";
import LateralBar from "@/components/LateralBart/LateralBar";
import TeslaCar from "@/public/tesla-car.svg";
import Car from "@/types/car";
import cars from "@/data/cars";
import DashboardStation from "@/components/DashboardStations/DashboardStation";

const Dashboard: NextPage = () => {

    const [selectedCar, setSelectedCar] = useState<Car>({
        name: "Tesla",
        model: "Model X",
        image: TeslaCar,
        battery: 89,
        range: 230,
        capacity: 80,
    });
    
    const [dashboradPage, setDashboardPage] = useState("stations");

    return (
        <div className="flex justify-center">
            <div className="flex w-[1800px] h-screen overflow-hidden">
                <LateralBar selectedCar={selectedCar} onSelectCar={setSelectedCar} cars={cars} dashboradPage={dashboradPage} setDashboardPage={setDashboardPage} />
                {dashboradPage === "stations" && (
                    <DashboardStation selectedCar={selectedCar}/> 
                )}
            </div>
        </div>
    );
};

export default Dashboard;
