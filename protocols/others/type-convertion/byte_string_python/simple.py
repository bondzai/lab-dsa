string1 = "a"
string1_binary = ''.join(format(ord(i), '08b') for i in string1)
string1_bytea = string1.encode()

print('==========')
print(string1_binary)
print(string1_bytea)