/* Alta3 Research | Indresh
   SOLUTION 02 - struct Virtmach and two methods
   ipset(string) - string is the IP to set on the struct 
   diskexpand(int) - int is the number of GB by which to expand the disk
   display() - display information about the struct      */

   package main

   import ("fmt")


   type Virtmach struct {

    ip       string
    hostname string
    diskgb   int
    ram      int

   }


   func (v *Virtmach) ipset(ip string) {
        v.ip = ip

   }


   func (v *Virtmach)diskexpand(gb int) {
        v.diskgb  = v.diskgb + gb

   }


   func (v Virtmach)display() {
	   fmt.Println("Primary IP Address:", v.ip) 
	   fmt.Println("Hostname:", v.hostname)
	   fmt.Println("Disk GB:", v.diskgb)
	   fmt.Println("RAM:", v.ram)

   }

   func main() {

      vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8}
      vm1.display()
      vm1.ipset("192.168.77.33")
      vm1.diskexpand(3)
      vm1.display()


   }

