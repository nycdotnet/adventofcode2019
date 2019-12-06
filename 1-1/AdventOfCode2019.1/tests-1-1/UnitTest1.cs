using app1_1;
using System;
using Xunit;

namespace tests_1_1
{
    public class UnitTest1
    {
        [Theory]
        [InlineData(12, 2)]
        [InlineData(14, 2)]
        [InlineData(1969, 654)]
        [InlineData(100756, 33583)]
        public void FuelToLaunchIsCorrect(int mass, int expectedFuelRequired)
        {
            Assert.Equal(expectedFuelRequired, FuelCalculator.FuelRequiredToLaunchMass(mass));
        }

        [Theory]
        [InlineData(1, 0)]
        [InlineData(14, 2)]
        [InlineData(1969, 966)]
        [InlineData(100756, 50346)]
        public void FuelToLaunchIsCorrectIncludingExtraFuel(int mass, int expectedFuelRequired)
        {
            Assert.Equal(expectedFuelRequired, FuelCalculator.FuelRequiredToLaunchMassIncludingTheFuel(mass));
        }
    }
}
