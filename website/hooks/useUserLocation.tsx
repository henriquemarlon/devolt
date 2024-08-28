import { useEffect, useState } from "react";

export default function useUserLocation() {
    const [userLocation, setUserLocation] = useState<[number, number]>([-22.98368, -43.21224]);
    const [city, setCity] = useState<string>("");

    useEffect(() => {
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition((position) => {
            const { latitude, longitude } = position.coords;
            setUserLocation([latitude, longitude]); 
            fetchCityName(latitude, longitude);
            });
        }
        }, []);
    
        const fetchCityName = async (latitude: number, longitude: number) => {
            const response = await fetch(
                `https://nominatim.openstreetmap.org/reverse?format=json&lat=${latitude}&lon=${longitude}&addressdetails=1`
            );
            const data = await response.json();
        
            if (data && data.address) {
                setCity(data.address.city || data.address.town || data.address.village || "Unknown city");
            } else {
                console.error("Error fetching city name");
            }
        };
        
    return [userLocation, city];
}