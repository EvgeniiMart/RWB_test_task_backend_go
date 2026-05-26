import json
from datetime import datetime

from nats.aio.client import Client as NATS
import asyncio


async def main():
    nc = NATS()

    await nc.connect("nats://nats:4222")

    events = [
        {
            "query": "apple",
            "delta": 10,
            "timestamp": datetime.utcnow().isoformat() + "Z"
        },
        {
            "query": "banana"
        }
    ]

    await nc.publish(
        "events",
        json.dumps(events).encode()
    )

    print("events sent")

    await nc.close()


asyncio.run(main())