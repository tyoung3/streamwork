(M edit.Model) <- 3(MV edit.ModelView) ; 
(MV)           <- 3(C edit.Controller) ; 
(C)            <- 1(V edit.View)       ; 

(V) <- 2(C);
(C)1 <- 2(MV);
(MV)1 <- 1(M); 

