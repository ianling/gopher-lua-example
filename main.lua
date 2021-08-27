my_dog = dog.new("Dingus")
print(my_dog:name())

my_dog:name("Bingus")
print(my_dog:name())

-- setting the function used when the dog speaks
function print_dog_sound ()
    print("Woof!")
end
my_dog:speak(print_dog_sound)

print("asking the dog to speak from Lua...")
my_dog:speak()()