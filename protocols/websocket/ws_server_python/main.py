import json
import asyncio
import websockets

async def handle_connection_1(websocket, path):
    while True:
        try:
            json_data = json.loads(await websocket.recv())
            print(f"Received JSON data from endpoint 1: {json_data}")
            await websocket.send(json.dumps(json_data))
        except json.decoder.JSONDecodeError as e:
            print(f"Invalid JSON: {e}")
            # await websocket.close(code=1000, reason='Invalid JSON')

async def handle_connection_2(websocket, path):
    while True:
        try:
            json_data = json.loads(await websocket.recv())
            print(f"Received JSON data from endpoint 2: {json_data}")
            await websocket.send(json.dumps(json_data))
        except json.decoder.JSONDecodeError as e:
            print(f"Invalid JSON: {e}")
            await websocket.close(code=1000, reason='Invalid JSON')

async def ws_router(websocket, path):
    if path == '/ws/1/':
        await handle_connection_1(websocket, path)
    elif path == '/ws/2/':
        await handle_connection_2(websocket, path)
    else:
        print("Invalid path")
        await websocket.close()

start_server = websockets.serve(ws_router, "localhost", 8000)

asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
