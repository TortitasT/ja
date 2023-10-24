call "build"

local arguments = args()
for k, v in pairs(arguments) do
  print(k, v)
end
