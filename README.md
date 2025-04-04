# pkt

`pkt` is a network toolkit built in Go; A practical set of tools for tracing, probing, and inspecting networks.

The first command is a custom `traceroute`, using TTL expiry and ICMP responses to uncover the route to a host. A `--watch` flag allows continuous tracing to help spot changes over time.

Commands follow a clean, script-friendly format:

```
pkt <command> <target> [flags]
```

Future functionality — like ping, packet sniffing, route inspection, and basic performance checks — will follow the same philosophy: focused, inspectable, and made for people who want to understand what’s really going on.

This project is being developed as a way to learn more about networking by building real tools from the ground up. If it’s useful to others along the way, all the better.

---

Licensed under [Apache 2.0](LICENSE).

Built with simplicity and clarity in mind.