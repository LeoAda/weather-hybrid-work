from config import Config

config = Config()
config.init()
#Menu
while True:  
    print("\nMAIN MENU")  
    print("1. Generate ")  
    print("2. Parameter")  
    print("5. Exit")  
    choice = int(input("Enter the Choice:"))  
    if choice == 5:
        break  
      
    else:  
        print("Oops! Incorrect Choice.")  