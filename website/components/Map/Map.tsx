import { useEffect, useState } from "react";
import L from "leaflet";
import "leaflet-defaulticon-compatibility";
import "leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css";
import "leaflet/dist/leaflet.css";
import { AttributionControl, MapContainer, MapContainerProps, Marker, Popup, TileLayer, useMap } from "react-leaflet";
import { CheckCircle } from "lucide-react";
import MapProps from "@/types/map";
import Station from "@/types/station";

//icons
const iAmHereIcon = L.icon({
  iconUrl: "/mapIcon.svg",
  iconRetinaUrl: "/mapIcon.svg",
  iconSize: [32, 32],
  popupAnchor: [-1, -16],
});

const stationIcon = L.icon({
  iconUrl: "/mapIcon2.svg",
  iconRetinaUrl: "/mapIcon2.svg",
  shadowSize: [80, 80],
  shadowAnchor: [30, 45],
  iconSize: [50, 50],
  iconAnchor: [25, 50],
  popupAnchor: [-1, -58],
});

// map updater
const MapUpdater = ({ mapCenter }: any) => {
  const map = useMap();

  useEffect(() => {
    map.flyTo(mapCenter, 14, {
      animate: true,
      duration: 5.0,
    });
  }, [mapCenter, map]);

  return null;
};

const Map = ({
  stations,
  width,
  height,
  center,
  userLocation,
  buttonText,
  roundedTopCorners,
  roundedBottomCorners,
  setSelectedStation,
}: MapProps) => {

  // map style
  const [containerStyle, setContainerStyle] = useState<MapContainerProps["style"]>({
    width: width || "100%",
    height: height || "550px",
    borderTopLeftRadius: roundedTopCorners ? "12px" : "0",
    borderTopRightRadius: roundedTopCorners ? "12px" : "0",
    borderBottomLeftRadius: roundedBottomCorners ? "12px" : "0",
    borderBottomRightRadius: roundedBottomCorners ? "12px" : "0",
    margin: "auto",
  });

  useEffect(() => {
    setContainerStyle({
      width: width || "100%",
      height: height || "550px",
      borderTopLeftRadius: roundedTopCorners ? "12px" : "0",
      borderTopRightRadius: roundedTopCorners ? "12px" : "0",
      borderBottomLeftRadius: roundedBottomCorners ? "12px" : "0",
      borderBottomRightRadius: roundedBottomCorners ? "12px" : "0",
      margin: "auto",
    });
  }, [width, height]);

  return (
    <>    
      <MapContainer
        center={center}
        zoom={1}
        attributionControl={false}
        scrollWheelZoom={true}
        style={containerStyle}
      >
        <AttributionControl prefix={false} position="bottomright" />

        <MapUpdater mapCenter={center} />

        <TileLayer
          url="https://tiles.stadiamaps.com/tiles/alidade_smooth_dark/{z}/{x}/{y}{r}.png"
          accessToken="3649afdf-ff6e-40b4-8d98-ef0deb099145"
        />

        <Marker icon={iAmHereIcon} position={userLocation}>
          <Popup><p className="text-white">You are here!</p></Popup>
        </Marker>

        {stations.map((station: Station, index: number) => {
          return (
            <Marker
              key={station.id}
              position={[station.latitude, station.longitude]}
              icon={stationIcon}
            >
              <Popup>
                <div className="leading-[1px] text-white">
                  <p className="text-center my-10 max-w-screen-md text-base">
                    {station.address || "Unnamed station"}{" "}
                  </p>
                  <div className="flex my-0 py-0 gap-2 justify-center items-center">
                    <CheckCircle size={20} color="#86ffb8" className="" />
                    <p className="font-bold text-base text-green-300">
                      Compatible
                    </p>
                  </div>
                  <div className="bg-neutral-800 rounded-lg max-w-max mx-auto mb-2 shadow px-4 flex">
                    <p className="text-lg font-medium text-center w-full">
                      {station.price_per_credit}$ / $VOLT 
                    </p>
                  </div>

                  {buttonText && (
                    <button
                      onClick={() => {
                        if (setSelectedStation) {
                          setSelectedStation(station);
                        }
                      }}
                      className="bg-[#3aff4e] text-base w-full mt-4 py-1 rounded-lg text-[#1e1e1e] font-bold hover:bg-green-400 transition"
                    >
                      {buttonText}
                    </button>
                  )}
                </div>
              </Popup>
            </Marker>
          );
        })}
      </MapContainer>

      <style>
        {`
        .leaflet-popup-content-wrapper, .leaflet-popup-tip {
          background-color: #222;
          border-radius: 10px;
          text-color: white;
          margin: 0;
        }
        .leaflet-popup-content p {
          margin: 10px 0;
          font-family: __Poppins_35a7f6, sans-serif;
      }
        `}
      </style>

    </>
  );
};

export default Map;
