Moore's Law:
------------

It predicted that Transistor density would double every 2 years. Although this isn't a physical law, it is just an observation.

Why Moore's Law is no longer valid:
-----------------------------------

-	**Power/Temperature Problem**: Transistors consume power when switching. Higher the number of Transistors, higher the power consumption. Higher the power consumption, higher temperature. At greater temperatures the physical chip would melt.

-	**Dynamic Power**: `P = α * CFV^2` Here, Swing Voltage V is a crucial factor as it's square is directly proportional to Power P. Thus Voltage should be as low as possible.

-	**Dennard Scaling**: "Voltage should scale with Transistor size". This causes problems as if the voltage level goes below a certain threshold voltage, noise problems occur. Additionally, this does not consider leakage power.

So how do we get better hardware?
---------------------------------

**Multi-Core Systems**: As in the Dynamic Power equation, Frequency F cannot be increased, multiple additional cores are added. This requires the employing of parallelism. Tasks are made to run on multiple cores rather than attempting to make a more powerful core.
