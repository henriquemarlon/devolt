export default async function useFetchAddress(latitude: number, longitude: number): Promise<string | null > {
    const response = await fetch(
        `https://nominatim.openstreetmap.org/reverse?format=json&lat=${latitude}&lon=${longitude}&addressdetails=1`
    );
    const data = await response.json();
    if (data && data.address) {
        const neighborhood = data.address.neighbourhood || data.address.suburb || data.address.village || data.address.town || data.address.city;
        return `${data.address.road} ${data.address.house_number}, ${neighborhood}, ${data.address.city}`;
    } else {
        console.error("Error fetching city name");
        return null;
    }
}