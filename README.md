# PacketDeck

Docker Desktop extension for visual container network inspection and packet capture.

## Why

Docker Desktop shows container names, images, and ports — but nothing about what's actually happening on the network between containers. Debugging multi-container networking currently requires CLI tools, tcpdump, and Wireshark. PacketDeck brings network visibility natively into Docker Desktop.

## Features (Planned)

- **Tier 1 (MVP)**: Interactive network topology map — which containers are on which networks, port mappings, connections
- **Tier 2**: Live connection monitor — TCP/UDP states, byte counts per container
- **Tier 3**: Packet capture — select container, capture traffic, filter, export pcap
- **Tier 4**: Application-layer inspection — HTTP viewer, DNS logging

## Status

Pre-development scaffold. See [HANDOFF.md](HANDOFF.md) for next steps.

## Development

```bash
make build-extension    # Build the extension image
make install-extension  # Install into Docker Desktop
make update-extension   # Update after changes
```

## License

[MIT](LICENSE) © 2026 Herb Hall
