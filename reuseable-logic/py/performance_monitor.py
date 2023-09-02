import time

def performance_monitor(func):
    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        elapsed_time = end_time - start_time
        print(f"Function '{func.__name__}' took {elapsed_time:.6f} seconds to execute.")
        return result
    return wrapper

# Example usage
@performance_monitor
def some_function():
    # Simulate some time-consuming task
    time.sleep(2)
    print("Task completed")

some_function()
