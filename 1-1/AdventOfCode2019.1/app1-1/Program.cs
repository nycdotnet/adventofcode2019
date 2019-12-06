using System;
using System.IO;

namespace app1_1
{
    class Program
    {
        static void Main(string[] args)
        {
            var moduleMassInputFile = File.ReadAllText(args[0]).Replace("\r\n","\n").Split("\n", StringSplitOptions.RemoveEmptyEntries);
            var totalFuelRequired = 0;
            foreach (var moduleMass in moduleMassInputFile)
            {
                var fuelRequired = FuelCalculator.FuelRequiredToLaunchMassIncludingTheFuel(int.Parse(moduleMass));
                Console.WriteLine($"Fuel required to launch module of mass {moduleMass} *including the fuel* is {fuelRequired}.");
                totalFuelRequired += fuelRequired;
            }
            Console.WriteLine($"Total fuel required to launch all {moduleMassInputFile.Length} modules *plus the fuel* is {totalFuelRequired}.");
        }
    }

    public static class FuelCalculator
    {
        public static int FuelRequiredToLaunchMass(int mass) => (int)Math.Floor(mass / 3.0) - 2;

        public static int FuelRequiredToLaunchMassIncludingTheFuel(int mass)
        {
            var allFuelRequired = FuelRequiredToLaunchMass(mass);
            var lastFuelAdded = allFuelRequired;
            while (true)
            {
                var newFuelForFuel = FuelRequiredToLaunchMass(lastFuelAdded);
                if (newFuelForFuel <= 0)
                {
                    break;
                }
                allFuelRequired += newFuelForFuel;
                lastFuelAdded = newFuelForFuel;
            }
            return Math.Max(0, allFuelRequired);
        }
    }
}
