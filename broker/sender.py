import asyncio
import json
from datetime import datetime, timezone

from nats.aio.client import Client as NATS


async def send_event(nc, event):
    await asyncio.sleep(event["start"])

    while True:
        payload = [
            {
                "query": event["query"],
                "delta": event["delta"],
                "timestamp": datetime.now(timezone.utc).isoformat()
            }
        ]

        await nc.publish(
            "events",
            json.dumps(payload).encode()
        )

        if event["repeat"] == -1:
            break

        await asyncio.sleep(event["repeat"])


async def main():
    nc = NATS()

    await nc.connect("nats://nats:4222")

    with open("events.json") as f:
        events = json.load(f)

    tasks = []

    for event in events:
        tasks.append(
            asyncio.create_task(send_event(nc, event))
        )

    await asyncio.gather(*tasks)

    await nc.close()


asyncio.run(main())