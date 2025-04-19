# Sync

- Use channels when passing ownership of data
- Use mutexes for managing state

## go vet

use go vet to help you detected subtle bugs

# don't use embedding because it's convenient

- Think about the effect embedding has on your public API
- Do you really want to expose these methods and have people coupling their own code to them?
- for mutexes, this could be potentiallly disastrous in very unpredictable and weird ways
  some very strange bugs that will be hard to track down
