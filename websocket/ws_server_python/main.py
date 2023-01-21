import asyncio
import websockets

async def ws_router(websocket, path):
    if path == '/ws/1':
        await handle_connection(websocket, path)
    else:
        print("Invalid path")
        await websocket.close()

start_server = websockets.serve(ws_router, "localhost", 8000)

asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
