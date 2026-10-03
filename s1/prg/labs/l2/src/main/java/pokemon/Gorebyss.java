package pokemon;

import move.Acupressure;
import move.Barrier;
import move.RazorLeaf;
import move.TailWhip;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Type;

final public class Gorebyss extends Pokemon {
  public Gorebyss(String name, int level) {
    super(name, level);
    this.setType(Type.WATER);
    this.setStats(55, 84, 105, 114, 75, 52);
    this.setMove(
      new Acupressure(),
      new RazorLeaf(),
      new TailWhip(),
      new Barrier()
    );
  }  
}
