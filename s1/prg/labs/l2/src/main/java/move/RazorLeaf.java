package move;

import ru.ifmo.se.pokemon.PhysicalMove;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.Type;

final public class RazorLeaf extends PhysicalMove {
  public RazorLeaf() {
    super(Type.GRASS, 55, 0.95);
  }

  @Override
  public double calcCriticalHit(Pokemon p1, Pokemon p2) {
    if (3 * p1.getStat(Stat.SPEED) / 512 > Math.random()) {
      System.out.println("Critical hit!");
      return 2;
    } else {
      return 1;
    }
  }

  @Override
  public String describe() {
    return "uses ability \"Razor Leaf\"";
  }
}
