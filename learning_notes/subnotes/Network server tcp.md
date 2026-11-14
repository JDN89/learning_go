[[web stuff to look up]]

# Network Server TCP

Een poort openzetten betekent meestal dat een proces op een lokaal IP en poort luistert en dat firewall/NAT-regels inkomende traffic naar die poort toelaten. Dat is niet hetzelfde als "outbound en inbound communiceren".

## Inbound vs outbound
- inbound: een remote host maakt verbinding met deze server.
- outbound: deze server maakt verbinding met een remote host.

## IP en poorten
- Een IP adres identificeert een host of interface.
- Een poort identificeert de service die op die host luistert.
- Een server kan meerdere poorten gebruiken voor verschillende diensten.

## TCP
- TCP is connection-oriented, betrouwbaar en levert data in order af.
- HTTP draait meestal boven TCP; het applicatieprotocol bepaalt de semantiek.

## Trust
- `mTLS` gebruikt certificaten om client en server te authenticeren op transportniveau.
- `JWT` is applicatielaag-identiteit en claims.
- Ze lossen verschillende problemen op en kunnen samen gebruikt worden.

## Related
- [[Heraclesx MTLS security and certificaten|mTLS security and certificates]]
- [[Web security]]
- [[API auth]]
