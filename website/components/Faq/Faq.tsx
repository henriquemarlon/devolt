import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "../../components/ui/accordion"

export function AccordionDemo() {
    return (
        <div className="flex justify-center flex-col items-center my-24 mx-4 md:mx-0">
            <p className="text-2xl md:text-5xl font-semibold md:px-8 mb-2">FAQ</p>
            <div className="mt-4 md:mt-12 max-w-screen-2xl w-[90%] px-4">
                <Accordion type="single" collapsible className="w-full">
                    <AccordionItem value="item-1">
                        <AccordionTrigger className="text-lg md:text-xl">How do I start selling energy with DeVolt?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            To start selling energy, register your interest on our platform. After meeting our specifications, install our software at your station to connect with buyers and efficiently sell energy.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-2">
                        <AccordionTrigger className="text-lg md:text-xl">What are the benefits of using DeVolt for EV charging?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            DeVolt provides accessible, cost-effective, and eco-friendly charging options. Our decentralized network ensures easy access to charging stations powered by renewable energy.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-3">
                        <AccordionTrigger className="text-lg md:text-xl">Is the DeVolt platform secure?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            Yes, DeVolt prioritizes security with advanced encryption and blockchain technology to secure transactions and ensure transparency, providing a safe and reliable service.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-4">
                        <AccordionTrigger className="text-lg md:text-xl">What advantages does DeVolt offer over traditional energy solutions?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            DeVolt leverages Web3 technologies to create a transparent, efficient, and user-governed energy marketplace. This reduces overhead costs, eliminates intermediaries, and provides real-time data integrity and pricing transparency.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-5">
                        <AccordionTrigger className="text-lg md:text-xl">Can anyone install a DeVolt charging station?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            Yes, anyone who meets DeVolt&apos;s technical and safety standards can install a charging station. Our platform facilitates the process, making it straightforward to join as a provider.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-6">
                        <AccordionTrigger className="text-lg md:text-xl">How does DeVolt handle energy pricing?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            Energy pricing on DeVolt is determined by market dynamics. Sellers set their prices based on supply and demand, ensuring competitive and fair pricing for all users.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-7">
                        <AccordionTrigger className="text-lg md:text-xl">What types of energy can be sold on DeVolt?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            DeVolt supports the sale of various types of renewable energy, including solar, wind, and hydro, depending on the capabilities and resources of the energy producer.
                        </AccordionContent>
                    </AccordionItem>
                    <AccordionItem value="item-10">
                        <AccordionTrigger className="text-lg md:text-xl">What is required to maintain a DeVolt charging station?</AccordionTrigger>
                        <AccordionContent className="text-neutral-400 md:text-xl">
                            Maintenance requirements are minimal, focusing primarily on safety checks and software updates to ensure optimal performance and compliance with DeVolt standards.
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
            </div>
        </div>
    )
}
